package models

import (
	"context"
	"time"

	
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"github.com/hmprgm/financial-planner/db"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Hash     string             `bson:"hash"`
	Email    string             `bson:"email"`
	Token    string             `bson:"token,omitempty"`
	// RefreshToken string             `bson:"refresh_token,omitempty"`
	CreatedAt primitive.DateTime `bson:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
}

type UserModel struct {
	collection *mongo.Collection
}

func (m *UserModel) Create(username, password, email string) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := &User{
		ID:        primitive.NewObjectID(),
		Username:  username,
		Hash:      string(hash),
		Email:     email,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	_, err := m.collection.InsertOne(context.TODO(), user)
	return err
}

func (m *UserModel) Login(username, password string) (string, error) {
	var user User
	err := m.collection.FindOne(context.TODO(), map[string]any{"username": username}).Decode(&user)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password))
	if err != nil {
		return "", err
	}

	token, err := db.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	// Store the token in the user's document
	update := map[string]any{
		"$set": map[string]any{
			"token":      token,
			"updated_at": primitive.NewDateTimeFromTime(time.Now()),
		},
	}
	_, err = m.collection.UpdateByID(context.TODO(), user.ID, update)
	if err != nil {
		return "", err
	}

	return token, nil
}

