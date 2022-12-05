package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitConnection() (string, error) {
	runtime := os.Getenv("RUNTIME_ENV")
	if runtime == "" {
		return "Can not get RUNTIME_ENV, aborting...", nil
	}
	log.Printf("RUNTIME_ENV %v", runtime)
	if err := godotenv.Load(".env." + runtime); err != nil {
		// log.Println("No .env file found")
		return "No .env file found", err
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		// log.Fatal("You must set your 'MONGODB_URI' environmental variable")
		return "You must set your 'MONGODB_URI' environmental variable", nil
	}
	log.Printf("Connection String: %v", uri)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		// panic(err)
		return "can not connect to mongodb", err
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			// panic(err)
			log.Fatal("can not disconnect to mongodb")
		}
	}()

	return "", nil
}
