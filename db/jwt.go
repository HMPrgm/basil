package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("fake_secret") // TODO Change this to a secure key

func GenerateJWT(userID primitive.ObjectID) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID.Hex(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
}

func GetUserIDFromToken(token *jwt.Token) (primitive.ObjectID, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return primitive.ObjectID{}, jwt.ErrSignatureInvalid
	}

	userIDHex, ok := claims["user_id"].(string)
	if !ok {
		return primitive.ObjectID{}, jwt.ErrSignatureInvalid
	}

	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return userID, nil
}