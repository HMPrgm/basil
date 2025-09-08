package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hmprgm/financial-planner/db/models"
)

type application struct {
	config config
	store models.Storage
}

type config struct {
	addr     string
	frontend string
}

func (app *application) mount() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{app.config.frontend},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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

	expenses := protected.Group("/expenses")

	expenses.POST("/", app.createExpense)
	expenses.GET("/", app.getExpenses)
	expenses.GET("/:id", app.getExpenseByID)
	expenses.PUT("/:id", app.updateExpense)
	expenses.DELETE("/:id", app.deleteExpense)

	return r
}

func (app *application) run(r *gin.Engine) error {
	if err := r.Run(app.config.addr); err != nil {
		return err
	}

	log.Printf("Server running on %s", app.config.addr)
	return nil
}