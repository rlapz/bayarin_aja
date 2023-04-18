package delivery

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rlapz/bayarin_aja/config"
	"github.com/rlapz/bayarin_aja/controller"
	"github.com/rlapz/bayarin_aja/middleware"
	"github.com/rlapz/bayarin_aja/repo/psql"
	"github.com/rlapz/bayarin_aja/usecase"

	_ "github.com/lib/pq"
)

type RestApi struct {
	http   http.Server
	db     *sql.DB
	engine *gin.Engine
	config *config.App
}

func NewRestApiDelivery() RestApi {
	conf := config.NewAppConfig()
	engine := gin.Default()

	return RestApi{
		http: http.Server{
			Addr:    conf.Api.HostPort,
			Handler: engine,
		},
		engine: engine,
		config: conf,
	}
}

func (self *RestApi) initDb() error {
	var err error
	dbCfg := self.config.Db
	dbParams := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		dbCfg.Username, dbCfg.Password, dbCfg.HostPort,
		dbCfg.Name, dbCfg.SslMode,
	)

	self.db, err = sql.Open("postgres", dbParams)
	if err != nil {
		return err
	}

	// test db connection
	return self.db.Ping()
}

func (self *RestApi) Run() error {
	if err := self.initDb(); err != nil {
		return err
	}

	// register router version 1
	self.v1()

	// spawning a new go routine for the http server
	go func(ctx *RestApi) {
		if err := ctx.http.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatalln(err.Error())
			}
		}
	}(self)

	// waiting for interrupt signal
	self.waitSignal()
	return nil
}

// routes
// version: 1
func (self *RestApi) v1() {
	rg := self.engine.Group("/v1")

	custRepo := psql.NewPsqlCustomerRepo(self.db)
	custActivityRepo := psql.NewPsqlCustomerActivityRepo(self.db)
	paymentRepo := psql.NewPsqlPaymentRepo(self.db)
	tokenRepo := psql.NewPsqlTokenRepo(self.db)

	tokenUsecase := usecase.NewTokenUsecase(tokenRepo)
	custActivityUsecase := usecase.NewCustomerActivityUsecase(custActivityRepo)
	custUsecase := usecase.NewCustomerUsecase(custRepo, custActivityUsecase, tokenUsecase)
	paymentUsecase := usecase.NewPaymentUsecase(paymentRepo)

	midTokenValidator := middleware.NewTokenValidator(tokenUsecase)
	midFunc := midTokenValidator.TokenValidate(self.config.Secret.Key)

	controller.NewCustomerController(
		rg,
		custUsecase,
		custActivityUsecase,
		midFunc,
		&self.config.Secret,
	)

	controller.NewPaymentController(
		rg,
		paymentUsecase,
		midFunc,
		&self.config.Secret,
	)
}

// helper
// gracefully stop the server
func (self *RestApi) waitSignal() {
	moeChan := make(chan os.Signal, 1)

	signal.Notify(moeChan, os.Interrupt, syscall.SIGTERM)
	<-moeChan

	// Add timeout 3 seconds to make sure all background process stopped
	timeout := 3 * time.Second
	fmt.Println()

	log.Println("Received a signal")
	log.Printf("Add timeout %d seconds to make sure all background process stopped cleanly\n", timeout/time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := self.http.Shutdown(ctx); err != nil {
		log.Fatalln("http.Shutdown:", err.Error())
	}

	<-ctx.Done()

	log.Println("The RESt API has been closed gracefully")
}
