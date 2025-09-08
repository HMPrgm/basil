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
	Expenses interface {
		Create(userID primitive.ObjectID, amount float64, category string, description string) error
		GetByUserID(userID primitive.ObjectID) ([]*ExpenseOutput, error)
		GetByCategory(userID primitive.ObjectID, category string) ([]*ExpenseOutput, error)
		GetByID(expenseID primitive.ObjectID) (*ExpenseOutput, error)
		Update(expenseID primitive.ObjectID, amount float64, category string, description string) error
		Delete(expenseID primitive.ObjectID) error
	}

}

func NewMongoStore(client *mongo.Client) Storage {
	return Storage{
		Users: &UserModel{
			collection: client.Database("main").Collection("users"),
		},
	}
}
