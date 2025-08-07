package main

import (
	"curp-scraper/internal/curp"
	"curp-scraper/internal/health"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not loaded: %v", err)
	}

	router := http.NewServeMux()

	// Set up routes
	router.Handle("/curp/", http.StripPrefix("/curp/", curp.Handler()))
	router.Handle("/health", health.Handler())

	// Get port from environment variable, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Starting CURP API server on port %s", port)

	log.Fatal(http.ListenAndServe(":"+port, router))

}
