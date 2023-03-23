package expense

import (
	exp "expense-logger/web/app/models/expense"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ////////// //
// Parameters //
// ////////// //

type latestParam struct {
	Count  int `json:"cnt"`
	Offset int `json:"offset"`
}

type getByDateParam struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

// //////// //
// Payloads //
// //////// //

type ItemPayload struct {
	Name  string   `json:"name" bson:"name"`
	Price uint32   `json:"price" bson:"price"`
	Count uint     `json:"count" bson:"count"`
	Unit  exp.Unit `json:"unit" bson:"unit"`
}

func (payload *ItemPayload) parse() exp.Item {
	var item = exp.Item{
		Name:  payload.Name,
		Price: payload.Price,
		Count: payload.Count,
		Unit:  payload.Unit,
	}

	return item
}

// Total in cent
type ExpensePayload struct {
	DateTime time.Time     `json:"datetime" bson:"datetime"`
	ShopName string        `json:"shopName" bson:"shopName"`
	Total    uint32        `json:"total" bson:"total"`
	Items    []ItemPayload `json:"items" bson:"items"`
}

func (payload *ExpensePayload) parse(resItem *mongo.InsertManyResult) exp.Expense {
	expense := exp.Expense{
		DateTime: payload.DateTime,
		ShopName: payload.ShopName,
		Total:    payload.Total,
	}

	expense.ItemIDs = make([]primitive.ObjectID, 0, len(resItem.InsertedIDs))
	for _, id := range resItem.InsertedIDs {
		expense.ItemIDs = append(expense.ItemIDs, id.(primitive.ObjectID))
	}

	return expense
}
