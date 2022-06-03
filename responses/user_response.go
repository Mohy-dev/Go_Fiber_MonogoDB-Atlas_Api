package responses

import "github.com/gofiber/fiber/v2"

// UserResponse is a struct that contains the json format scheme response of the user
type UserResponse struct {
	Status  int        `json:"status"` // Struct tags
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"` // *fiber.Map is empty interface that welcomed to append and hold any type of data
}
