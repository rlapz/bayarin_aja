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

	<-moeChan

	log.Println("The RESt API has been closd gracefully")
}
