package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", SlashRoute)

	http.ListenAndServe(":4321", r)
}

func SlashRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}
