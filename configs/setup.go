package configs

import (
	"context" // Context is a way to signal cancellation of a goroutine.
	"fmt"     // fmt is the standard Go package for formatted I/O.
	"time"    // time is the standard Go package for time functions.

	"go.mongodb.org/mongo-driver/mongo"         // mongo is the MongoDB driver.
	"go.mongodb.org/mongo-driver/mongo/options" // options is the MongoDB driver options.
)

func ConnectDB() *mongo.Client {
	// Set client options and connect to the server with the options provided in the options.ClientOptions, The context is used to specify the options for the connection.
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		panic(err) // Panic if there is an error
	}

	// Create a context to cancel the connection if the program is interrupted
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Cancel the context when the function returns
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	// Check the connection by pinging the server
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	// Print the connection status
	fmt.Println("Connected to MongoDB!")
	// return the client to the main function to use it
	return client
}

// Initialize the mongoDB client instance
var DB *mongo.Client = ConnectDB()

// Initialize the mongoDB database instance and collection instance within that database
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	// Get the collection from the database, just change the database name "golang-fiber-api"
	return client.Database("golang-fiber-api").Collection(collectionName)
}
