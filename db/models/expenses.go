package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"context"
	"time"
)

type Expense struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	UserID primitive.ObjectID `bson:"user_id"`
	Amount float64 `bson:"amount"`
	Category string `bson:"category"`
	Date primitive.DateTime `bson:"date"`
	Description string `bson:"description"`
	CreatedAt primitive.DateTime `bson:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
}

type ExpenseOutput struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Amount      float64            `bson:"amount" json:"amount"`
	Category    string             `bson:"category" json:"category"`
	Date       primitive.DateTime  `bson:"date" json:"date"`
	Description string             `bson:"description" json:"description"`
}

type ExpenseModel struct {
	collection *mongo.Collection
}

func (m *ExpenseModel) Create(userID primitive.ObjectID, amount float64, category, description string) error {
	expense := &Expense{
		ID:         primitive.NewObjectID(),
		UserID:     userID,
		Amount:     amount,
		Category:   category,
		Description: description,
		CreatedAt:  primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:  primitive.NewDateTimeFromTime(time.Now()),
	}
	_, err := m.collection.InsertOne(context.TODO(), expense)
	return err
}

func (m *ExpenseModel) GetByUserID(userID primitive.ObjectID) ([]*ExpenseOutput, error) {
	filter := map[string]any{"user_id": userID}
	cursor, err := m.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var expenses []*ExpenseOutput
	if err := cursor.All(context.TODO(), &expenses); err != nil {
		return nil, err
	}
	return expenses, nil
}

func (m *ExpenseModel) GetByID(expenseID primitive.ObjectID) (*ExpenseOutput, error) {
	var expense ExpenseOutput
	err := m.collection.FindOne(context.TODO(), map[string]any{"_id": expenseID}).Decode(&expense)
	if err != nil {
		return nil, err
	}
	return &expense, nil
}

func (m *ExpenseModel) GetByCategory(userID primitive.ObjectID, category string) ([]*ExpenseOutput, error) {
	filter := map[string]any{"user_id": userID, "category": category}
	cursor, err := m.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var expenses []*ExpenseOutput
	if err := cursor.All(context.TODO(), &expenses); err != nil {
		return nil, err
	}
	return expenses, nil
}

func (m *ExpenseModel) Update(expenseID primitive.ObjectID, amount float64, category, description string) error {
	update := map[string]any{
		"$set": map[string]any{
			"amount":      amount,
			"category":    category,
			"description": description,
			"updated_at":  primitive.NewDateTimeFromTime(time.Now()),
		},
	}
	_, err := m.collection.UpdateByID(context.TODO(), expenseID, update)
	return err
}

func (m *ExpenseModel) Delete(expenseID primitive.ObjectID) error {
	_, err := m.collection.DeleteOne(context.TODO(), map[string]any{"_id": expenseID})
	return err
}