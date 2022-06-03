package models

import (
	"time" // Time package

	"go.mongodb.org/mongo-driver/bson/primitive" // MongoDB primitive package
)

// User is a struct that contains the user data
type User struct {
	Id        primitive.ObjectID `json:"id omitempty"`
	Name      string             `json:"name,omitempty,validate:"required""`
	Age       int                `json:"age",omitempty,validate:"required"`
	Email     string             `json:"email", omitempty, validate:"required"`
	Gender    string             `json:"gender",omitempty, default:"x"`
	Location  Location           `json:"location", omitempty, default:"Empty"`
	Title     string             `json:"title", omitempty, validate:"required"`
	CreatedAt time.Time          `json:"created_at", omitempty`
	UpdatedAt time.Time          `json:"updated_at", omitempty`
}

type Location struct {
	Street    string    `json:"street,omitempty"`
	City      string    `json:"city,omitempty"`
	State     string    `json:"state,omitempty"`
	VisitedAt time.Time `json:"visitedAt,omitempty"`
}
