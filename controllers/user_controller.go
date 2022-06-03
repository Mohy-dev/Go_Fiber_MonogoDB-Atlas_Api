package controllers

import (
	"context"                   // used to cancel the operation when the function returns or when the timeout expires
	"fiber-mongo-api/configs"   // import the configs
	"fiber-mongo-api/models"    // import the models
	"fiber-mongo-api/responses" // import the responses
	"net/http"                  // import the http package
	"net/mail"                  // import the mail package
	"time"                      // import the time package

	"github.com/go-playground/validator/v10"     // import the validator package
	"github.com/gofiber/fiber/v2"                // import the fiber package
	"go.mongodb.org/mongo-driver/bson"           // import the bson package
	"go.mongodb.org/mongo-driver/bson/primitive" // import the bson package
	"go.mongodb.org/mongo-driver/mongo"          // import the mongo package
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users") // Collection name here is "users"
var validate = validator.New()                                                    // Initialize the validator

// Create user
func CreateUser(c *fiber.Ctx) error {
	// context is used to cancel the operation when the function returns or when the timeout expires (10 seconds)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel() // Cancel the context when the function returns

	// validate the request body  and return the error if any
	if err := c.BodyParser(&user); err != nil {
		return response(c, "error", http.StatusBadRequest, err.Error())
	}

	// validate the required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return response(c, "error", http.StatusBadRequest, validationErr.Error())
	}

	currentTime := time.Now()
	// Insert the user data to mongo object to insert it later
	newUser := models.User{
		Id:        primitive.NewObjectID(),
		Name:      user.Name,
		Age:       user.Age,
		Email:     user.Email,
		Gender:    user.Gender,
		Location:  user.Location,
		Title:     user.Title,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	if !validEmail(newUser.Email) {
		return response(c, "error", http.StatusBadRequest, "Invalid Email")
	}

	// insert the new user into the database and return the error if any
	result, err := userCollection.InsertOne(ctx, newUser) // insert one user object mongo operation
	if err != nil {
		return response(c, "error", http.StatusInternalServerError, err.Error())
	}

	// return the user data to the client with the status code 201 Created and the response body as JSON format with the user data
	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

// Get user by id
func GetUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel() // Cancel the context when the function returns

	// get the user id from the request
	id := c.Params("id")

	// ObjectIdHex is a helper function that converts a hexadecimal string representation of an ObjectId into a primitive.ObjectID.
	objId, _ := primitive.ObjectIDFromHex(id)

	// find the user by the id and return the error if any.
	// Mongo read the data as BSON format and decode it to the user object.
	err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user) // FindOne operation to find the user by the id
	if err != nil {
		// return the error to the client with the status code 404 Not Found and the response body as JSON format with the error
		return response(c, "error", http.StatusInternalServerError, err.Error())
	}

	// return the user data to the client with the status code 200 OK and the response body as JSON format with the user data
	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": user}})
}

// func to get all users
func GetAllUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []models.User
	defer cancel()

	iter, err := userCollection.Find(ctx, bson.M{}) // Find operation to find all users
	if err != nil {
		return response(c, "error", http.StatusInternalServerError, err.Error())
	}

	// iterate through the users and append them to the users array
	defer iter.Close(ctx)
	for iter.Next(ctx) {
		var user models.User
		if err := iter.Decode(&user); err != nil {
			return response(c, "error", http.StatusInternalServerError, err.Error())
		}
		users = append(users, user)
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": users}})
}

// Edit user
func EditUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")
	var user models.User
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	if err := c.BodyParser(&user); err != nil {
		return response(c, "error", http.StatusBadRequest, err.Error())
	}

	if !validEmail(user.Email) {
		return response(c, "error", http.StatusBadRequest, "Invalid Email")
	}

	if validatorErr := validate.Struct(&user); validatorErr != nil {
		return response(c, "error", http.StatusBadRequest, validatorErr.Error())
	}

	update := bson.M{"name": user.Name, "title": user.Title, "age": user.Age, "email": user.Email, "location": user.Location, "gender": user.Gender, "updatedAt": time.Now()}

	// update the user by the id and return the error if any
	result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

	if err != nil {
		return response(c, "error", http.StatusInternalServerError, err.Error())
	}

	var updatedUser models.User
	if result.MatchedCount == 1 {
		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)

		if err != nil {
			return response(c, "error", http.StatusInternalServerError, err.Error())
		}
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedUser}})
}

// Delete user
func DeleteUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	// delete the user by the id and return the error if any
	result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId}) // DeleteOne operation to delete the user by the id
	if err != nil {
		return response(c, "error", http.StatusInternalServerError, err.Error())
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "User ID not found!"}})
	}

	// return the user data to the client with the status code 200 OK and the response body as JSON format with the user data
	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": result}})
}

// Delete all users
func PurgeUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// delete all users and return the error if any
	result, err := userCollection.DeleteMany(ctx, bson.M{}) // DeleteMany operation to delete all users
	if err != nil {
		return response(c, "error", http.StatusInternalServerError, err.Error())
	}

	// return the user data to the client with the status code 200 OK and the response body as JSON format with the user data
	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": result}})
}

// Helper function to match the email format
func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// generic response function to return the response to the client
func response(c *fiber.Ctx, message string, status int, data interface{}) error {
	return c.Status(status).JSON(responses.UserResponse{Status: status, Message: message, Data: &fiber.Map{"data": data}})
}
