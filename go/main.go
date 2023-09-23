package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", root)

	http.ListenAndServe(":4321", r)
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
