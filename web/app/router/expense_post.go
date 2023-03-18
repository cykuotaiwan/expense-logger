package router

import (
	"encoding/json"
	exp "expense-logger/web/app/models/expense"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

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

func PostExpense(c *gin.Context) {
	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	var payload ExpensePayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	log.Print(payload)
	log.Print(len(payload.Items))

	// insert items
	items := make([]exp.Item, len(payload.Items))
	for _, it := range payload.Items {
		items = append(items, it.parse())
	}
	log.Print(len(items))
	log.Print(items)
	resItem, err := exp.InsertItem(items)
	log.Print(err)
	log.Print(resItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	if len(resItem.InsertedIDs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "item count shouldn't be zero",
		})
		return
	}

	// insert expense
	expense := payload.parse(resItem)
	resExp, err := exp.InsertExpense(&expense)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	// success response
	insertedItemLen := len(resItem.InsertedIDs)
	insertedExpenseLen := 0
	if resExp != nil {
		insertedExpenseLen = 1
	}
	c.JSON(http.StatusOK, gin.H{
		"itemInsertedCnt":    insertedItemLen,
		"expenseInsertedCnt": insertedExpenseLen,
	})
}
