package main

import (
	"fiber-mongo-api/configs"
	"fiber-mongo-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello World"}) // &fiber.Map = &map[string]interface{}
	})

	// Run the database connection
	configs.ConnectDB()

	// Routes
	routes.UserRoute(app)

	// Run the server
	app.Listen(":6000")
}

// MONGOURI=mongodb+srv://golang:12345@cluster0.m1nvx.mongodb.net/?retryWrites=true&w=majority
