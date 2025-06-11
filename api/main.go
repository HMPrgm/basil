package main

import (
	"log"
	"net/http"
	"github.com/hmprgm/financial-planner/internal/env"
	"github.com/joho/godotenv"
)

const version = "0.0.1"

func setupDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	setupDotEnv()
	addr := env.GetString("ADDR", ":8080")

	srv := http.Server{
		Addr: addr,
	}

	log.Fatal(srv.ListenAndServe())
	log.Printf("Server started on %s", addr)
}
