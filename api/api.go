package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hmprgm/financial-planner/db/models"
)

type application struct {
	config config
	store models.Storage
}

type config struct {
	addr string
}

func (app *application) mount() *gin.Engine {
	r := gin.Default()

	// Public routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	// User routes
	r.POST("/login", app.login)
	r.POST("/register", app.register)

	// Protected routes
	protected := r.Group("/api")
	protected.Use(JWTAuthMiddleware())
	protected.GET("/user", app.getUserInfo)

	return r
}

func (app *application) run(r *gin.Engine) error {
	if err := r.Run(app.config.addr); err != nil {
		return err
	}

	log.Printf("Server running on %s", app.config.addr)
	return nil
}