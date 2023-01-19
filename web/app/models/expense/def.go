package expense

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Item struct {
	ItemID uint32  `json:"itemId" bson:"itemId"`
	Name   string  `json:"name" bson:"name"`
	Price  float32 `json:"price" bson:"price"`
	Count  uint    `json:"count" bson:"count"`
}

type Expense struct {
	ExpenseID uint32    `json:"expenseId" bson:"expenseId"`
	DateTime  time.Time `json:"datetime" bson:"datetime"`
	ShopName  string    `json:"shopName" bson:"shopName"`
	Total     float32   `json:"total" bson:"total"`
	ItemSet   []Item    `json:"itemSet" bson:"itemSet"`
}

var expCollection *mongo.Collection
var itemCollection *mongo.Collection
var ctx context.Context

func SetCollections(client *mongo.Client, context *context.Context) {
	expCollection = client.Database("expense").Collection("expense")
	itemCollection = client.Database("expense").Collection("items")
	ctx = *context
}
