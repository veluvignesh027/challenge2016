package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/RealImage/challenge2016/internal/database"
	"github.com/RealImage/challenge2016/internal/server"
	"github.com/joho/godotenv"
	"github.com/veluvignesh027/log"
)

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Info("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Info("Server forced to shutdown with error: %v", err)
	}

	log.Info("Server exiting")

	done <- true
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := database.LoadCSVFile(); err != nil {
		log.Fatal("Error loading CSV file: ", err)
	}
}

func main() {
	flag.Parse()
	log.Info("Starting the Application...")
	server := server.NewServer()

	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(server, done)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(fmt.Sprintf("http server error: %s", err))
	}

	<-done
	log.Info("Graceful shutdown complete.")
}
