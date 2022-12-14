package main

import (
	"api-postgresql/configs"
	"api-postgresql/handlers"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(any(err))
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
