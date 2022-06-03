package routes

import (
	"fiber-mongo-api/controllers" // Import the controllers

	"github.com/gofiber/fiber/v2" // Fiber is a high performance web framework for Go
)

// UserRoute is a function that contains the routes of the user
func UserRoute(app *fiber.App) {
	app.Post("/user", controllers.CreateUser)       // Create a new user
	app.Get("/user/:id", controllers.GetUser)       // Get a user by id
	app.Get("/user", controllers.GetAllUsers)       // Get all users
	app.Put("/user/:id", controllers.EditUser)      // Edit a user
	app.Delete("/user/:id", controllers.DeleteUser) // Delete a user
	app.Delete("/users", controllers.PurgeUsers)    // Delete all users
}
