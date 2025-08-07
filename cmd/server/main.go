package main

import (
	"curp-scraper/internal/curp"
	"curp-scraper/internal/health"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	// Set up routes
	http.Handle("/curp", curp.Handler())
	http.Handle("/health", health.Handler())

	// Start server
	port := "8080"
	log.Printf("Starting CURP API server on port %s", port)
	log.Printf("CURP endpoint: http://localhost:%s/curp?curp=YOUR_CURP_HERE", port)
	log.Printf("Health check: http://localhost:%s/health", port)

	log.Fatal(http.ListenAndServe(":"+port, router))

}
