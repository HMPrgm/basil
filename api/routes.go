package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hmprgm/financial-planner/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *gin.Context) {
	cursor, err := mongoClient.Database("main").Collection("users").Find(context.TODO(), bson.D{{}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var users []bson.M
	if err = cursor.All(context.TODO(), &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func Login(c *gin.Context) {
    var req struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }
    
	var user models.User
	err := mongoClient.Database("main").Collection("users").FindOne(context.TODO(), bson.M{"username": req.Username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(req.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	token, _ := GenerateJWT(user.Username)
	user.Token = token
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Register(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user := models.CreateUser(req.Username, req.Password, req.Email, "", "")
	_, err := mongoClient.Database("main").Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func Mount(r *gin.Engine) {
	r.GET("/users", JWTAuthMiddleware(), GetUsers)
	r.POST("/login", Login)
	r.POST("/register", Register)
}