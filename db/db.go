package db

import (
	"context"
	"os"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
)



var uri string 
var mongoClient *mongo.Client

// The init function will run before our main function to establish a connection to MongoDB
func New() (*mongo.Client, error) {
	setupDotEnv()
	uri = os.Getenv("DB_CONNECTION_STRING")

	if err := connectToMongoDB(); err != nil {
		return nil, err
	}
	return mongoClient, nil
}

func setupDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env File")
	}
}

func connectToMongoDB() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	mongoClient = client
	return err
}