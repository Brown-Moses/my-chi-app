package handlers

import (
	"chi-app/internal/models"
	"chi-app/internal/services"
	"encoding/json"
	"net/http"
)

//handlers for every service..

// getstudent/view
func GetStudentHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Application/json")

	// Encode the struct to JSON and write to response
	err := json.NewEncoder(w).Encode(services.ViewStudent())
	if err != nil {
		panic(err)
	}

}

// add student
func AddStudentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var newStudent models.Student
	err := json.NewDecoder(r.Body).Decode(&newStudent)

	if err != nil {
		http.Error(w, "INVALID JSON", http.StatusBadRequest)
	}

	services.AddStudent(newStudent)

	err = json.NewEncoder(w).Encode(map[string]string{
		"message": "Student added successfully",
		"name":    newStudent.Name,
	})

}
