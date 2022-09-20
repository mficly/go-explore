package db

import(
	"log"
	"os"
	"context"
	"github.com/joho/godotenv"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitConnection(){
	runtime := os.Getenv("RUNTIME_ENV")
	if runtime == "" {
		panic("Can not get RUNTIME_ENV, aborting...")
	}
	log.Printf("RUNTIME_ENV %v", runtime)
	if err := godotenv.Load(".env."+ runtime); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable")
	}
	log.Printf("Connection String: %v", uri)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	
}

