package main

import (
	"github.com/atoyegbe/sre-bootcamp/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

func main() {
	app := fiber.New()
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/healthcheck",
	}))
	port := ":8000"
	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/student/:studentId", handlers.GetStudent)
	v1.Get("/students", handlers.GetAllStudents)
	v1.Post("/student", handlers.CreateStudent)
	v1.Put("/student/:studentId", handlers.UpdateStudent)
	v1.Delete("/student/:studentId", handlers.DeleteStudent)

	app.Listen(port)
}
