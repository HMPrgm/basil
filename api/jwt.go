package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"strings"
    "github.com/gin-gonic/gin"
    "net/http"
)

var jwtKey = []byte("fake_secret") // TODO Change this to a secure key

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
}

func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            c.Abort()
            return
        }
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        token, err := ValidateJWT(tokenString)
        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
        c.Next()
    }
}