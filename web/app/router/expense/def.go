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
	Id    primitive.ObjectID `json:"id" bson:"id"`
	Name  string             `json:"name" bson:"name"`
	Price uint32             `json:"price" bson:"price"`
	Count uint               `json:"count" bson:"count"`
	Unit  exp.Unit           `json:"unit" bson:"unit"`
}

func (payload *ItemPayload) parse() exp.Item {
	var item = exp.Item{
		Id:    payload.Id,
		Name:  payload.Name,
		Price: payload.Price,
		Count: payload.Count,
		Unit:  payload.Unit,
	}

	return item
}

type ExpensePayload struct {
	Id       primitive.ObjectID `json:"id" bson:"id"`
	DateTime time.Time          `json:"datetime" bson:"datetime"`
	ShopName string             `json:"shopName" bson:"shopName"`
	Total    uint32             `json:"total" bson:"total"` //cent
	Items    []ItemPayload      `json:"items" bson:"items"`
}

func (payload *ExpensePayload) parseWithInsertResult(resItems *mongo.InsertManyResult) exp.Expense {
	expense := exp.Expense{
		Id:       payload.Id,
		DateTime: payload.DateTime,
		ShopName: payload.ShopName,
		Total:    payload.Total,
	}

	expense.ItemIDs = make([]primitive.ObjectID, 0, len(resItems.InsertedIDs))
	for _, id := range resItems.InsertedIDs {
		expense.ItemIDs = append(expense.ItemIDs, id.(primitive.ObjectID))
	}

	return expense
}

func (payload *ExpensePayload) parse() exp.Expense {
	expense := exp.Expense{
		Id:       payload.Id,
		DateTime: payload.DateTime,
		ShopName: payload.ShopName,
		Total:    payload.Total,
	}

	expense.ItemIDs = make([]primitive.ObjectID, 0, len(payload.Items))
	for _, items := range payload.Items {
		expense.ItemIDs = append(expense.ItemIDs, items.Id)
	}

	return expense
}
