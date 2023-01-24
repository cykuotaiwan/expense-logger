package expense

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Unit uint

const (
	Uundefine Unit = iota
	Uea
	Ulb
)

// Price in cent
type Item struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Price     uint32             `json:"price" bson:"price"`
	Count     uint               `json:"count" bson:"count"`
	Unit      Unit               `json:"unit" bson:"unit"`
	ExpenseID primitive.ObjectID `json:"expenseId" bson:"expenseId"`
}

// Total in cent
type Expense struct {
	Id       primitive.ObjectID   `json:"_id,omitempty" bson:"_id"`
	DateTime time.Time            `json:"datetime" bson:"datetime"`
	ShopName string               `json:"shopName" bson:"shopName"`
	Total    uint32               `json:"total" bson:"total"`
	ItemIDs  []primitive.ObjectID `json:"itemIDs" bson:"itemIDs"`
}

var expCollection *mongo.Collection
var itemCollection *mongo.Collection
var ctx context.Context

func SetCollections(client *mongo.Client, context *context.Context) {
	if os.Getenv("OPMODEs") == "RELEASE" {
		expCollection = client.Database("expense").Collection("expense")
		itemCollection = client.Database("expense").Collection("items")
	} else {
		expCollection = client.Database("expense-debug").Collection("expense")
		itemCollection = client.Database("expense-debug").Collection("items")
	}

	ctx = *context
}

func (item Item) IsEmpty() bool {
	empty := false
	empty = empty || len(item.Name) == 0
	empty = empty || item.Price == 0
	empty = empty || item.Count == 0
	empty = empty || item.Unit == Uundefine

	return empty
}

func (exp Expense) IsEmpty() bool {
	empty := false
	empty = empty || exp.DateTime.IsZero()
	empty = empty || len(exp.ShopName) == 0
	empty = empty || exp.Total == 0

	return empty
}
