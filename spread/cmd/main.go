package main

import (
	"context"

	"log"
	"net/http"
	"os"
	"os/signal"

	"runtime/trace"
	"time"

	"github.com/kelvin950/spread/config"
	server "github.com/kelvin950/spread/internals/adapters"
	"github.com/kelvin950/spread/internals/core/api"
)

func main() {

	f, _ := os.Create("trace.out")
	trace.Start(f)

	defer trace.Stop()
	api := api.NewApiMock()

	log.Printf("%T ,%T", api, api.S3Client)
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
