package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)



func (app *application) createExpense(c *gin.Context) {
	var input struct {
		UserID      primitive.ObjectID `json:"user_id"`
		Amount      float64            `json:"amount"`
		Category    string             `json:"category"`
		Date        primitive.DateTime `json:"date"`
		Description string             `json:"description"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := app.store.Expenses.Create(input.UserID, input.Amount, input.Category, input.Description); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Expense created successfully"})
}

func (app *application) getExpenses(c *gin.Context) {
	userID := c.MustGet("userID").(primitive.ObjectID)
	category := c.Query("category")

	if category != "" {
		expenses, err := app.store.Expenses.GetByCategory(userID, category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, expenses)
		return
	}

	expenses, err := app.store.Expenses.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, expenses)
}

func (app *application) getExpenseByID(c *gin.Context) {
	expenseID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expense ID"})
		return
	}
	expense, err := app.store.Expenses.GetByID(expenseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, expense)
}