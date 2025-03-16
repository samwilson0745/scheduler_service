package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"scheduler_service/db"
	"scheduler_service/routes"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DB := db.ConnectDB()
	psGrs, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get db: %v", err)
	}

	router := routes.MainRouter()
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT)

	// run server in go routine

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()
	log.Printf("Serve started on :8080")

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	// Shutdown the server gracefully
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown failed: %v", err)
	}

	// Close the database connection
	if err := psGrs.Close(); err != nil {
		log.Printf("Failed to close database connection: %v", err)
	}

	log.Println("Server and database connection closed gracefully")
}
