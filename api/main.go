package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/login", Login)

	authorized := r.Group("/admin")
	authorized.Use(AuthMiddleware())
	{
		authorized.GET("/data", func(c *gin.Context) {
			username, _ := c.Get("username")
			c.JSON(http.StatusOK, gin.H{"message": "This is protected data for " + username.(string)})
		})
	}

	r.Run(":8080")
}

// login, logout, me handlers omitted for brevity
