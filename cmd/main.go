package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	route "github.com/jegj/linktly/internal/api/routes"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Heartbeat("/ping"))

	r.Mount("/test", route.PingRoutes())
	// r.Mount("/api/v1", route.PingRoutes())

	// Start the HTTP server
	log.Printf("starting server on %s\n", os.Getenv("SERVER_ADDRESS"))
	err := http.ListenAndServe(os.Getenv("SERVER_ADDRESS"), r)
	if err != nil {
		log.Fatal(err)
	}
}
