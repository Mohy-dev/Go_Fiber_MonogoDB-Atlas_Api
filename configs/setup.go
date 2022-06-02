package configs

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	// Set client options and connect to the server with the options provided in the options.ClientOptions, The context is used to specify the options for the connection.
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		panic(err)
	}

	// Create a context to cancel the connection if the program is interrupted
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
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

// Initialize the mongoDB database instance and collection instance
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("golang-fiber-api").Collection(collectionName)
}
