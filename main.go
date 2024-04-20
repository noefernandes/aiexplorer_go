package main

import (
	"aiexplorer/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

const BASE_URL = "/api/v1/aitool"

func main() {
	r := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"Content-Type", "Accept", "Authorization"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	r.Use(c.Handler)

	r.Get(BASE_URL, handlers.GetAll)
	r.Get(BASE_URL+"/{id}", handlers.Get)
	r.Post(BASE_URL, handlers.Save)
	r.Put(BASE_URL, handlers.Update)

	http.ListenAndServe(":9000", r)
}
