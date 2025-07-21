package main

import (
	"github.com/gin-gonic/gin"
    "context"
	"net/http"
	"strings"
    "github.com/hmprgm/financial-planner/db"
)

func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            c.Abort()
            return
        }
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        token, err := db.ValidateJWT(tokenString)
        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
        // Store the token in the context for later use
        c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "token", token))
        c.Next()
    }
}

