package expense_test

import (
	config "expense-logger/configs"
	db "expense-logger/web/app/models"
	exp "expense-logger/web/app/models/expense"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
		expense.Total = 0
		res, err := exp.UpdateExpense(expense, id)
		if err != nil {
			t.Error(err.Error())
		}
		if res == nil {
			t.Error("update fail")
		}
	})
}
