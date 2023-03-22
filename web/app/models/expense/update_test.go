package expense_test

import (
	config "expense-logger/configs"
	db "expense-logger/web/app/models"
	exp "expense-logger/web/app/models/expense"
	util "expense-logger/web/app/util"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUpdateItem(t *testing.T) {
	util.NewDBConnection()
	defer util.EndDBConnection()

	// generate test data
	var newItem = []exp.Item{
		{
			Name:  "Potatoes",
			Price: 249,
			Unit:  exp.Ulb,
			Count: 3,
		},
	}

	// insert test data
	resTest, _ := exp.InsertItem(newItem)

	t.Run("valid value", func(t *testing.T) {
		id := resTest.InsertedIDs[0].(primitive.ObjectID)
		item := newItem
		item[0].Id = id
		item[0].Count = 10

		res, err := exp.UpdateItem(item[0])
		assert.Nil(t, err)
		if err != nil {
			t.Error(err.Error())
		}
		if res == nil {
			t.Error("update fail")
		}
	})
}

func TestUpdateExpense(t *testing.T) {
	// db connection
	config.Init()
	db.Init()
	defer db.Close()

	// generate test data
	var newExpense = exp.Expense{
		DateTime: time.Now(),
		ShopName: "Loblaws",
		Total:    2011,
	}

	// insert test data
	resTest, _ := exp.InsertExpense(&newExpense)

	t.Run("valid value", func(t *testing.T) {
		id := (*resTest).InsertedID.(primitive.ObjectID)
		expense := newExpense
		expense.Id = id
		expense.Total = 0
		res, err := exp.UpdateExpense(expense)
		if err != nil {
			t.Error(err.Error())
		}
		if res == nil {
			t.Error("update fail")
		}
	})
}
