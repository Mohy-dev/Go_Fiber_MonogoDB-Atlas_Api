package responses

import "github.com/gofiber/fiber"

// UserResponse is a struct that contains the response of the user
type UserResponse struct {
	Status  int        `json:"status"` // Struct tag
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}
