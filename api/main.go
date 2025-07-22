package main

import (
	"log"
	"github.com/hmprgm/financial-planner/db"
	"github.com/hmprgm/financial-planner/db/models"
)

func main() {

	cfg := config{
		addr: ":8080",
	}
	
	mongoClient, err := db.New()
	if err != nil {
		log.Fatal("Could not connect to MongoDB: ", err)
	}
	log.Println("MongoDB connection established")

	store := models.NewMongoStore(mongoClient)
	app := &application{
		config: cfg,
		store:  store,
	}
	
	r := app.mount()

	log.Fatal(app.run(r))
}

