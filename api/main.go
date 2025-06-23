package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	

	"github.com/joho/godotenv"
)

var uri string 

var mongoClient *mongo.Client

// The init function will run before our main function to establish a connection to MongoDB
func init() {
	setupDotEnv()
	uri = os.Getenv("DB_CONNECTION_STRING")

	if err := connectToMongoDB(); err != nil {
		log.Fatal("Could not connect to MongoDB: ", err)
	}
}

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	// Mount the routes
	Mount(r)

	r.Run() 
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