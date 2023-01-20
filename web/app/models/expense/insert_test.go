package expense_test

import (
	config "expense-logger/configs"
	db "expense-logger/web/app/models"
	exp "expense-logger/web/app/models/expense"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertItem(t *testing.T) {
	// db connection
	config.Init()
	db.Init()
	defer db.Close()

	// setup test data
	var newItems = []exp.Item{
		{
			Name:  "Potatoes",
			Price: 249,
			Unit:  exp.Ulb,
			Count: 3,
		},
		{
			Name:  "Milk",
			Price: 1398,
			Unit:  exp.Uea,
			Count: 1,
		},
	}
	// main test
	t.Run("nil", func(t *testing.T) {
		res, err := exp.InsertItem(nil)
		if err != nil || res != nil {
			t.Errorf(err.Error())
		}
	})
	t.Run("empty array", func(t *testing.T) {
		res, err := exp.InsertItem([]exp.Item{})
		if err != nil || res != nil {
			t.Errorf(err.Error())
		}
	})
	t.Run("empty items", func(t *testing.T) {
		res, err := exp.InsertItem([]exp.Item{{}, {}})
		if err != nil || res != nil {
			t.Errorf(err.Error())
		}
	})
	// test insert items
	t.Run("two items", func(t *testing.T) {
		res, err := exp.InsertItem(newItems)
		if err != nil || res == nil {
			t.Errorf(err.Error())
		}
	})

}

func TestInsertExpense(t *testing.T) {
	// db connection
	config.Init()
	db.Init()
	defer db.Close()

	// setup test data
	var newExpense = exp.Expense{
		DateTime: time.Now(),
		ShopName: "Loblaws",
		Total:    2011,
	}

	// main test
	t.Run("nil", func(t *testing.T) {
		res, err := exp.InsertExpense(nil)
		if err != nil || res != nil {
			t.Errorf(err.Error())
		}
	})
	t.Run("empty", func(t *testing.T) {
		res, err := exp.InsertExpense(&exp.Expense{})
		if err != nil || res != nil {
			t.Errorf(err.Error())
		}
	})
	t.Run("valid value", func(t *testing.T) {
		res, err := exp.InsertExpense(&newExpense)
		if err != nil || res == nil {
			t.Errorf(err.Error())
		}
	})
}

func TestInsertItemExpense(t *testing.T) {
	// db connection
	config.Init()
	db.Init()
	defer db.Close()

	// setup test data
	var newItems = []exp.Item{
		{
			Name:  "Potatoes",
			Price: 249,
			Unit:  exp.Ulb,
			Count: 3,
		},
		{
			Name:  "Milk",
			Price: 1398,
			Unit:  exp.Uea,
			Count: 1,
		},
	}
	var newExpense = exp.Expense{
		DateTime: time.Now(),
		ShopName: "Loblaws",
		Total:    2011,
	}

	t.Run("both item and expense", func(t *testing.T) {
		resItem, errItem := exp.InsertItem(newItems)
		if errItem != nil || resItem == nil {
			t.Errorf(errItem.Error())
		}

		newExpense.ItemIDs = make([]primitive.ObjectID, 0, len(resItem.InsertedIDs))
		for _, id := range resItem.InsertedIDs {
			newExpense.ItemIDs = append(newExpense.ItemIDs, id.(primitive.ObjectID))
		}
		resExp, errExp := exp.InsertExpense(&newExpense)
		if errExp != nil || resExp == nil {
			t.Errorf(errExp.Error())
		}
	})
}
