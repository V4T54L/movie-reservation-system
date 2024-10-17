package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/V4T54L/movie-reservation-system/internals/config"
	"github.com/V4T54L/movie-reservation-system/internals/routes"
)

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	<-ctx.Done()
	log.Println("shutting down gracefully, press CTRL+C again to force")

	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()

	if err := apiServer.Shutdown(ctx2); err != nil {
		log.Printf("Server forced to shutdown with error : %s", err)
	}

	log.Println("server exiting...")
	done <- true
}

func main() {
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", config.GetConfig().ServerPort),
		Handler:      routes.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	done := make(chan bool, 1)

	go gracefulShutdown(&server, done)

	log.Println("Starting server...")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	<-done

	log.Println("graceful shutdown complete")
}
