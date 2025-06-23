package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

func Mount(r *gin.Engine) {
	r.GET("/users", GetUsers)
}