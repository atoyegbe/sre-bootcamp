package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/atoyegbe/sre-bootcamp/database"
	"github.com/atoyegbe/sre-bootcamp/handlers"
	"github.com/atoyegbe/sre-bootcamp/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupApp() *fiber.App {
    app := fiber.New()

    // Initialize the database for testing
    database.DB, _ = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
    database.DB.AutoMigrate(&models.Student{})

	app.Get("/api/v1/student/:studentId", handlers.GetStudent)
	app.Get("/api/v1/students", handlers.GetAllStudents)
	app.Post("/api/v1/student", handlers.CreateStudent)
	app.Put("/api/v1/student/:studentId", handlers.UpdateStudent)
	app.Delete("api/v1/student/:studentId", handlers.DeleteStudent)

	app.Listen(":8000")
    return app
}


func TestGetAllStudents(t *testing.T) {
	app := setupApp()

	students := []models.Student{
		{Name: "John Doe", Age: 23, Email: "john.doe@example.com"},
		{Name: "Jane Smith", Age: 22, Email: "jane.smith@example.com"},
	}
	for _, s := range students {
		database.DB.Create(&s)
	}

	req := httptest.NewRequest("GET", "/api/v1/students", nil)
	resp, err := app.Test(req)

	require.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Parse response body
	var responseStudents []models.Student
	err = json.NewDecoder(resp.Body).Decode(&responseStudents)
	require.NoError(t, err)

	assert.Len(t, responseStudents, len(students))
	for i, s := range students {
		assert.Equal(t, s.Name, responseStudents[i].Name)
		assert.Equal(t, s.Age, responseStudents[i].Age)
		assert.Equal(t, s.Email, responseStudents[i].Email)
	}

}

func TestCreateStudent(t *testing.T) {
	app := setupApp()

	// Prepare student data
	student := models.Student{
		Name:  "New Student",
		Age:   20,
		Email: "new.student@example.com",
	}
	studentJSON, _ := json.Marshal(student)

	req := httptest.NewRequest("POST", "/api/v1/student", bytes.NewBuffer(studentJSON))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)


	require.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

	var createdStudent models.Student
	err = json.NewDecoder(resp.Body).Decode(&createdStudent)
	require.NoError(t, err)

	assert.Equal(t, student.Name, createdStudent.Name)
	assert.Equal(t, student.Age, createdStudent.Age)
	assert.Equal(t, student.Email, createdStudent.Email)
	assert.NotZero(t, createdStudent.ID)

}

func TestUpdateStudentDetails(t *testing.T) {
	app := setupApp()

	student := models.Student{
		Name:  "mike",
		Age:   25,
		Email: "original.email@example.com",
	}
	database.DB.Create(&student)

	updatedStudent := models.Student{
		Name:  "mike",
		Age:   26,
		Email: "mike.email@example.com",
	}
	updatedStudentJSON, _ := json.Marshal(updatedStudent)
	req := httptest.NewRequest("PUT", "/api/v1/student/"+ strconv.Itoa(int(student.ID)), bytes.NewBuffer(updatedStudentJSON))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	require.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var returnedStudent models.Student
	err = json.NewDecoder(resp.Body).Decode(&returnedStudent)
	require.NoError(t, err)

	assert.Equal(t, updatedStudent.Name, returnedStudent.Name)
	assert.Equal(t, updatedStudent.Age, returnedStudent.Age)
	assert.Equal(t, updatedStudent.Email, returnedStudent.Email)
	assert.Equal(t, student.ID, returnedStudent.ID)
}

func TestDeleteStudent(t *testing.T) {
	app := setupApp()

	student := models.Student{
		Name:  "deleteduser",
		Age:   30,
		Email: "deleteduser@example.com",
	}
	database.DB.Create(&student)

	req := httptest.NewRequest("DELETE", "/api/v1/student/"+ strconv.Itoa(int(student.ID)), nil)
	resp, err := app.Test(req)

	require.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var deletedStudent models.Student
	result := database.DB.First(&deletedStudent, student.ID)
	assert.Error(t, result.Error) // Should return an error as the student should not be found

}