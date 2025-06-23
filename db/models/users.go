package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string             `bson:"username"`
	Hash         string             `bson:"hash"`
	Email        string             `bson:"email"`
	Token        string             `bson:"token,omitempty"`
	RefreshToken string             `bson:"refresh_token,omitempty"`
	CreatedAt    primitive.DateTime `bson:"created_at"`
	UpdatedAt    primitive.DateTime `bson:"updated_at"`
}

func CreateUser(username, password, email string, token, refreshToken string) *User {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return &User{
		ID:           primitive.NewObjectID(),
		Username:     username,
		Hash:         string(hash),
		Email:        email,
		Token:        token,
		RefreshToken: refreshToken,
		CreatedAt:    primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:    primitive.NewDateTimeFromTime(time.Now()),
	}
}
