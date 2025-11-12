package main

import (
	"chi-app/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// saving some draft students in the list of Student struct... Incase we intergrate DB later> w shall the storage conn instead of a list arr

func main() {
	r := chi.NewRouter()

	r.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("System under maintainance..."))
	})

	// displaying a struct json to browser
	r.Get("/student", handlers.GetStudentHandler)

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		panic(err)
	}

}
