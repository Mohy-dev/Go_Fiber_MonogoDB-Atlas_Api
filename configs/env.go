package configs

import (
	// log is the standard Go package for logging.
	"os" // os is the Go standard library's interface to the operating system, covering things like file path

	"github.com/joho/godotenv" // godotenv is a Go library for loading environment variables from .env files
)

// EnvMongoURI is the mongoDB URI checker
func EnvMongoURI() string {
	err := godotenv.Load() // Loads all the env variables from .env file
	if err != nil {
		panic("Error loading .env file") // Panic if there is an error loading the .env file
	}
	return os.Getenv("MONGOURI") // Return the mongoDB URI
}
