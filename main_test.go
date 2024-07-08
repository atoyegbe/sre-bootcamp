package main

import (
	"net/http/httptest"
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

    // Create a test student
    student := models.Student{
        Name:  "John Doe",
        Age:   23,
        Email: "john.doe@example.com",
    }
    database.DB.Create(&student)

    // Make a GET request to the /students endpoint
    req := httptest.NewRequest("GET", "/api/v1/students", nil)
    resp, err := app.Test(req)

    require.NoError(t, err)
    assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestGetAllStudentsDetails(t *testing.T) {
	// Test the endpoint that gets details of all students
}

func TestCreateStudent(t *testing.T) {
	// Test the endpoint that creates a new student
}

func TestUpdateStudentDetails(t *testing.T) {
	// Test the endpoint that updates a student's details
}

func TestDeleteStudent(t *testing.T) {
	// Test the endpoint that deletes a student
}