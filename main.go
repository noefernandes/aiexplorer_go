package main

import (
	"aiexplorer/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

func main() {
	r := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedHeaders:   []string{"Content-Type", "Accept", "Origin", "Authorization"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	r.Use(c.Handler)

	r.Get("/", handlers.GetAll)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(":9000", r)
}
