package main

import (
	"project/database"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	panic("unimplemented")
}

func main() {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
