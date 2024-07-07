package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Student struct {
	Id       string `json:"id,omitempty"`
	Age      int64  `json:"age"`
	Name     string `json:"name"`
	Position int64  `json:"position"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Student{})

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

	v1.Get("/student/:studentId", func(c *fiber.Ctx) error {
		var student Student
		result := db.First(&student, c.Params("studentId"))
		return c.JSON(result)
	})

	v1.Get("/students", func(c *fiber.Ctx) error {
		var students []Student
		result := db.Find(&students)
		return c.JSON(result)
	})

	v1.Post("/student", func(c *fiber.Ctx) error {
		c.Accepts("json", "text")
		p := new(Student)
		if err := c.BodyParser(p); err != nil {
			return err
		}
		return c.JSON(p)
	})

	v1.Put("/student/:studentId", func(c *fiber.Ctx) error {
		student := new(Student)
		if err := c.BodyParser(student); err != nil {
			return err
		}
		db.Model(&student).Updates(student)
		return c.JSON(student)
	})

	v1.Delete("/student/:studentId", func(c *fiber.Ctx) error {
		db.Delete(&Student{}, c.Params("studentId"))
		return c.JSON(fiber.Map{
			"message": "Deleted!",
		})
	})

	app.Listen(port)
}
