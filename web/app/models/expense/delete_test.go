package expense_test

import (
	exp "expense-logger/web/app/models/expense"
	util "expense-logger/web/app/util"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDeleteItem(t *testing.T) {
	util.NewDBConnection()
	defer util.EndDBConnection()

	// generate test data
	var newItem = []exp.Item{
		{
			Name:  "TestDeletePotatoes",
			Price: 249,
			Unit:  exp.Ulb,
			Count: 3,
		},
		{
			Name:  "TestDeleteMilk",
			Price: 1398,
			Unit:  exp.Uea,
			Count: 1,
		},
	}
	var newExpense = exp.Expense{
		DateTime: time.Now(),
		ShopName: "TestDeleteLoblaws",
		Total:    2011,
	}

	var itemIds []primitive.ObjectID
	resItem, err := exp.InsertItem(newItem)
	if err == nil {
		for _, elem := range (*resItem).InsertedIDs {
			itemIds = append(itemIds, elem.(primitive.ObjectID))
		}
		newExpense.ItemIDs = itemIds
	}
	resExp, _ := exp.InsertExpense(&newExpense)
	expenseId := (*resExp).InsertedID.(primitive.ObjectID)
	t.Run("valid data", func(t *testing.T) {
		resDel, err := exp.DeleteItem(itemIds[0])

		assert.Nil(t, err)
		assert.Equal(t, int64(1), (*resDel).DeletedCount)

		resExp, _ := exp.GetExpenseLatest(1, 0)
		assert.Equal(t, expenseId.String(), resExp[0].Id.String())

		resItem, _ := exp.GetItemLatest(2, 0)
		assert.NotEqual(t, itemIds[0].String(), resItem[0].Id.String())
		assert.Equal(t, itemIds[1].String(), resItem[0].Id.String())
	})
	t.Run("delete non exist item", func(t *testing.T) {
		nonExistItemId := itemIds[0]
		nonExistItemId[0] -= 128

		resDel, err := exp.DeleteItem(nonExistItemId)
		assert.Nil(t, err)
		assert.Equal(t, int64(0), resDel.DeletedCount)
	})
}

func TestDeleteExpense(t *testing.T) {
	util.NewDBConnection()
	defer util.EndDBConnection()

	var newExpense = exp.Expense{
		DateTime: time.Now(),
		ShopName: "TestDeleteLoblaws",
		Total:    2087,
	}
	resExp, _ := exp.InsertExpense(&newExpense)
	expenseId := (*resExp).InsertedID.(primitive.ObjectID)

	t.Run("delete non exist item", func(t *testing.T) {
		nonExistExpenseId := expenseId
		nonExistExpenseId[0] -= 128

		resDel, err := exp.DeleteExpense(nonExistExpenseId)

		assert.Nil(t, err)
		assert.Equal(t, int64(0), resDel.DeletedCount)
	})

	t.Run("valid data", func(t *testing.T) {
		resDel, err := exp.DeleteExpense(expenseId)

		assert.Nil(t, err)
		assert.Equal(t, int64(1), (*resDel).DeletedCount)
	})
}
