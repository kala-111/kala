package routes

import (
	Db "project/database"
	"project/models"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/employee", ListEmployee)
	app.Get("/employee/:id", GetEmployeebyId)
	app.Post("/employee", CreateEmployee)
	app.Delete("/employee/:id", DeleteEmployee)
	app.Put("/employee/:id", UpdateEmployee)

}
func ListEmployee(c *fiber.Ctx) error { //fucntion fo ListEmployeee
	employees := []models.Employee{}
	Db.DB.Db.Find(&employees)
	return c.Status(200).JSON(employees)

}
func GetEmployeebyId(c *fiber.Ctx) error { //function of getby id
	id := c.Params("id")

	var employee models.Employee

	result := Db.DB.Db.Find(&employee, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(employee)
}
func CreateEmployee(c *fiber.Ctx) error {
	var employee models.Employee

	if err := c.BodyParser(employee); err != nil {
		return c.SendStatus(400)
	}
	Db.DB.Db.Create(&employee)
	return c.Status(201).JSON(employee)
}
func DeleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	var employee models.Employee
	if result := Db.DB.Db.Find(&employee, id); result.Error != nil {
		return c.SendStatus(404)
	}
	Db.DB.Db.Delete(&employee)
	return c.Status(200).JSON(&employee)
}
func UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	var employee models.Employee
	Db.DB.Db.Find(&employee, "id=? ", id)

	if employee.ID == ' ' {
		return c.Status(404).JSON(fiber.Map{"status": "Status not found"})
	}
	var UpdateEmployee models.Employee

	err := c.BodyParser(&UpdateEmployee)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "Status not found"})

	}
	UpdateEmployee.ID = employee.ID
	Db.DB.Db.Model(&employee).Updates(UpdateEmployee)

	Db.DB.Db.Save(&employee)
	return c.JSON(fiber.Map{"data": employee})

}
