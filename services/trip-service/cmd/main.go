package main

import (
	"log"
	"net/http"
	h "ride-sharing/services/trip-service/internal/infrastructure/http"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
)

func main() {
	log.Println("Starting Trip Service")
	inmemRepo := repository.NewInmemRepository()
	svc := service.NewTripService(inmemRepo)
	httpHandler := h.HttpHandler{Service: svc}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /preview", httpHandler.HandleTripPreview)

	server := &http.Server{
		Addr:    ":8083",
		Handler: mux,
	}

	log.Println("Starting Trip Service 2")

	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP server error: %v", err)
	}
}
