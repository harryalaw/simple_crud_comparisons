package main

import (
	"net/http"
    "encoding/json"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", root)
    r.Get("/json", getJson)

	http.ListenAndServe(":4321", r)
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

type Data struct {
    Id int32 
    Message string 
    Array []string
}

func getJson(w http.ResponseWriter, r *http.Request) {
    data := Data{
        Id: 1234,
        Message: "I'm a JSON value",
        Array: []string{ "here's", "some", "values", "in", "an", "array", "to", "make", "the", "object", "bigger"},
    }

    json, err := json.Marshal(data)
    if err != nil {
        w.Write([]byte("Oopsies"))
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(json)
}
