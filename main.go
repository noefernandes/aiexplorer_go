package main

import (
	"aiexplorer/configs"
	"aiexplorer/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

const BASE_URL = "/api/v1/"

func main() {
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"Content-Type", "Accept", "Authorization"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	r.Use(c.Handler)

	r.Get(BASE_URL+"aitools", handlers.GetAll)
	r.Get(BASE_URL+"aitool/{id}", handlers.Get)
	r.Post(BASE_URL+"aitool", handlers.Save)
	r.Put(BASE_URL+"aitool", handlers.Update)
	r.Delete(BASE_URL+"aitool", handlers.Delete)

	r.Get(BASE_URL+"tags", handlers.GetAllTags)
	r.Get(BASE_URL+"tag/{id}", handlers.GetTag)
	r.Post(BASE_URL+"tag", handlers.SaveTag)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
