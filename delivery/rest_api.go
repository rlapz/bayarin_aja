package delivery

import (
	"context"
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
	"github.com/rlapz/bayarin_aja/repo/json_repo"
	"github.com/rlapz/bayarin_aja/usecase"
)

type RestApi struct {
	http   http.Server
	engine *gin.Engine
	config *config.App
}

func NewRestApiDelivery() RestApi {
	conf := config.NewAppConfig()
	engine := gin.Default()

	return RestApi{
		http: http.Server{
			Addr: fmt.Sprintf(
				"%s:%s",
				conf.Api.Host,
				conf.Api.Port,
			),
			Handler: engine,
		},
		engine: engine,
		config: conf,
	}
}

func (self *RestApi) Run() error {
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

	json_db_path := self.config.DbJSONPath
	customerRepo := json_repo.NewJSONCustomerRepo(json_db_path)
	tokenRepo := json_repo.NewJSONTokenRepo(json_db_path)

	customerUsecase := usecase.NewCustomerUsecase(customerRepo, tokenRepo)
	tokenUsecase := usecase.NewTokenUsecase(tokenRepo)

	controller.NewCustomerController(
		rg,
		customerUsecase,
		tokenUsecase,
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
