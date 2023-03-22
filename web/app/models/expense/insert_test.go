package expense_test

import (
	exp "expense-logger/web/app/models/expense"
	util "expense-logger/web/app/util"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertItem(t *testing.T) {
	util.NewDBConnection()
	defer util.EndDBConnection()

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

		assert.Nil(t, err)
		assert.Nil(t, res)
	})
	t.Run("empty array", func(t *testing.T) {
		res, err := exp.InsertItem([]exp.Item{})

		assert.Nil(t, err)
		assert.Nil(t, res)
	})
	t.Run("empty items", func(t *testing.T) {
		res, err := exp.InsertItem([]exp.Item{{}, {}})

		assert.Nil(t, err)
		assert.Nil(t, res)
	})
	// test insert items
	t.Run("two items", func(t *testing.T) {
		res, err := exp.InsertItem(newItems)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Len(t, res.InsertedIDs, 2)
	})

}

func TestInsertExpense(t *testing.T) {
	util.NewDBConnection()
	defer util.EndDBConnection()

	// setup test data
	var newExpense = exp.Expense{
		DateTime: time.Now(),
		ShopName: "Loblaws",
		Total:    2011,
	}

	// main test
	t.Run("nil", func(t *testing.T) {
		res, err := exp.InsertExpense(nil)

		assert.Nil(t, err)
		assert.Nil(t, res)
	})
	t.Run("empty", func(t *testing.T) {
		res, err := exp.InsertExpense(&exp.Expense{})

		assert.Nil(t, err)
		assert.Nil(t, res)
	})
	t.Run("valid value", func(t *testing.T) {
		res, err := exp.InsertExpense(&newExpense)

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}

func TestInsertItemExpense(t *testing.T) {
	util.NewDBConnection()
	defer util.EndDBConnection()

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

		assert.Nil(t, errItem)
		assert.NotNil(t, resItem)
		assert.Len(t, resItem.InsertedIDs, 2)

		newExpense.ItemIDs = make([]primitive.ObjectID, 0, len(resItem.InsertedIDs))
		for _, id := range resItem.InsertedIDs {
			newExpense.ItemIDs = append(newExpense.ItemIDs, id.(primitive.ObjectID))
		}
		resExp, errExp := exp.InsertExpense(&newExpense)

		assert.Nil(t, errExp)
		assert.NotNil(t, resExp)
	})
}
