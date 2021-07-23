package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

var shutdownTimeout = flag.Duration("shutdown-timeout", 10*time.Second, "shutdown timeout (5s,5m,5h) before connections are cancelled")

const (
	service = "book-service"
)

func main() {
	r := mux.NewRouter()

	RegisterHandlers(r)

	serverPort := os.Getenv("PORT")

	srv := &http.Server{
		Addr:         ":" + serverPort,
		Handler:      r,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	gracefulShutdown(srv, serverPort)

}

func gracefulShutdown(srv *http.Server, serverPort string) {
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Printf("%s listening on %s:with %v timeout", service, serverPort, *shutdownTimeout)
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-stop

	log.Printf("%s shutting down ...\n", service)

	ctx, cancel := context.WithTimeout(context.Background(), *shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Printf("%s down\n", service)
}
