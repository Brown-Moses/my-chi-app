package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Student struct {
	ID    int    `json:"ID"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

func main() {
	r := chi.NewRouter()

	r.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("System under maintainance..."))
	})

	// displaying a struct json to browser
	r.Get("/student", func(w http.ResponseWriter, r *http.Request) {

		stud := Student{ID: 1, Name: "Mosesssss", Email: "Moses@gmail.com"}

		//set headers to tell the browser its json
		w.Header().Set("Content-Type", "Application/json")

		// Encode the struct to JSON and write to response
		json.NewEncoder(w).Encode(stud)

	})

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		panic(err)
	}

}
