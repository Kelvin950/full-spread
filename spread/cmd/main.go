package main

import (
	"context"
	"runtime"

	"log"
	"net/http"
	"os"
	"os/signal"

	"runtime/trace"
	"time"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/kelvin950/spread/config"
	server "github.com/kelvin950/spread/internals/adapters/httpServer"
	"github.com/kelvin950/spread/internals/core/api"
)

func main() {

	f, _ := os.Create("trace.out")
	trace.Start(f)
runtime.GOMAXPROCS(1)
	defer trace.Stop()

	cfg , err:= awscfg.LoadDefaultConfig(context.Background())

	if err!=nil{
		log.Fatal(err)
	}

	api := api.NewApi(cfg)


	server := server.NewServer(*api)
	conf := config.NewConfig()
	port := conf.GetKey("PORT")
	httpServer := http.Server{
		Addr:    port,
		Handler: server.Router,
	}

	errCh := make(chan error)
	go func() {

		log.Println("Starting server on port", port)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
		close(errCh)

	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	select {
	case <-sig:
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		log.Println("Shutting down server...")
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}

	case err := <-errCh:
		if err != nil {
			log.Fatal(err)
		}

	}

}
