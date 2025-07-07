package models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	Users interface {
		Create(username, password, email string) error
		Login(username, password string) (string, error)
	}
}

func NewMongoStore(client *mongo.Client) Storage {
	return Storage{
		Users: &UserModel{
			collection: client.Database("main").Collection("users"),
		},
	}
}
