package main

/*
- go get -u github.com/gofiber/fiber/v2 go.mongodb.org/mongo-driver/mongo github.com/joho/godotenv github.com/go-playground/validator/v10

- go get github.com/klauspost/compress
*/

import (
	"fiber-mongo-api/configs" // Import the configs
	"fiber-mongo-api/routes"  // import routes

	"github.com/gofiber/fiber/v2" // Fiber is a high performance web framework for Go
)

func main() {
	app := fiber.New() // Create a new Fiber instance

	// Json Message for that route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Welcome to Restful Api with mongoDB atlas"}) // &fiber.Map = &map[string]interface{}
	})

	// Run the database connection
	configs.ConnectDB()

	// Routes
	routes.UserRoute(app)

	// Run the server
	app.Listen(":6000")
}
