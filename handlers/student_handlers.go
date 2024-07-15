package handlers

import (
	"github.com/atoyegbe/sre-bootcamp/database"
	"github.com/atoyegbe/sre-bootcamp/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllStudents(c *fiber.Ctx) error {
	var students []models.Student
	result := database.DB.Find(&students)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}
	return c.JSON(students)
}

func GetStudent(c *fiber.Ctx) error {
	var student models.Student
	result := database.DB.First(&student, c.Params("studentId"))
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Student not found",
		})
	}
	return c.JSON(student)
}

func CreateStudent(c *fiber.Ctx) error {
	p := new(models.Student)
	if err := c.BodyParser(p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	newStudent := models.Student{Name: p.Name, Age: p.Age, Email: p.Email}
	res := database.DB.Create(&newStudent)
	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": res.Error.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(newStudent)
}

func UpdateStudent(c *fiber.Ctx) error {
	id := c.Params("studentId")
	var student models.Student
	if err := database.DB.First(&student, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Student not found",
		})
	}
	if err := c.BodyParser(&student); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	database.DB.Save(&student)
	return c.JSON(student)
}

func DeleteStudent(c *fiber.Ctx) error {
	id := c.Params("studentId")
	var student models.Student
	if err := database.DB.First(&student, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Student not found",
		})
	}
	database.DB.Delete(&student)
	return c.JSON(fiber.Map{
		"message": "Deleted!",
	})
}
