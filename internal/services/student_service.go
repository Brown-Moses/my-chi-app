package services

import "chi-app/internal/models"

var student = []models.Student{
	{ID: 1, Name: "Mosesssss", Email: "Moses@gmail.com"},
	{ID: 2, Name: "Wuruemmmm", Email: "Wuruem@gmail.com"},
	{ID: 3, Name: "Justinnnn", Email: "Justin@gmail.com"},
}

// adds a student
func AddStudent(s models.Student) error {
	student = append(student, s)
	return nil
}

//view/ return all students in the arr to browser
func ViewStudent() []models.Student {
	return student
}
