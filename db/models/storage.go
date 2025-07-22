package models

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Storage struct {
	Users interface {
		Create(username, password, email string) error
		Login(username, password string) (string, error)
		GetByID(userID primitive.ObjectID) (*User, error)
	}
}

func NewMongoStore(client *mongo.Client) Storage {
	return Storage{
		Users: &UserModel{
			collection: client.Database("main").Collection("users"),
		},
	}
}
