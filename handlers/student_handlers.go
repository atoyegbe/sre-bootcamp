package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/atoyegbe/sre-bootcamp/models"
	"github.com/atoyegbe/sre-bootcamp/database"
)


func GetAllStudents(c *fiber.Ctx) (error) {
	var students []models.Student
	result := database.DB.Find(&students)
	return c.JSON(result)
}

func GetStudent(c *fiber.Ctx) error {
	var student models.Student
	result := database.DB.First(&student, c.Params("studentId"))
	return c.JSON(result)
}

func CreateStudent(c *fiber.Ctx) error {
	c.Accepts("json", "text")
	p := new(models.Student)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	res := database.DB.Create(&p)
	return c.JSON(res)
}

func UpdateStudent(c *fiber.Ctx) error {
	student := new(models.Student)
	if err := c.BodyParser(student); err != nil {
		return err
	}
	database.DB.Model(&student).Updates(student)
	return c.JSON(student)
}

func DeleteStudent(c *fiber.Ctx) error {
	database.DB.Delete(&models.Student{}, c.Params("studentId"))
	return c.JSON(fiber.Map{
		"message": "Deleted!",
	})
}
