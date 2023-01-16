package expense

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Count uint    `json:"count"`
}

type expense struct {
	ID       string    `json:"id"`
	DateTime time.Time `json:"datetime"`
	ShopName string    `json:"shopName"`
	Total    float32   `json:"total"`
	ItemSet  []item    `json:"itemSet"`
}

var expCollection *mongo.Collection
var ctx context.Context

func SetExpenseCollection(client *mongo.Client, context *context.Context) {
	expCollection = client.Database("expense").Collection("expense")
	ctx = *context
}
