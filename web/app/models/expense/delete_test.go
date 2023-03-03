package expense_test

import (
	config "expense-logger/configs"
	db "expense-logger/web/app/models"
	exp "expense-logger/web/app/models/expense"
	"fmt"
	"strconv"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDeleteItem(t *testing.T) {
	// db connection
	config.Init()
	db.Init()
	defer db.Close()

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
		if err != nil {
			t.Error(err.Error())
		}
		if (*resDel).DeletedCount != 1 {
			t.Error(
				fmt.Errorf("deleted count should be 1, instead of %d",
					(*resDel).DeletedCount),
			)
		}

		resExp, _ := exp.GetExpenseLatest(1, 0)
		if resExp[0].Id.String() != expenseId.String() {
			t.Error(fmt.Errorf("expense id should be the same"))
		}

		resItem, _ := exp.GetItemLatest(2, 0)
		if resItem[0].Id.String() != itemIds[1].String() ||
			resItem[0].Id.String() == itemIds[0].String() {
			t.Error(fmt.Errorf("designated item is not deleted"))
		}
	})
	t.Run("delete non exist item", func(t *testing.T) {
		nonExistItemId := itemIds[0]
		nonExistItemId[0] -= 128

		resDel, err := exp.DeleteItem(nonExistItemId)
		if err != nil {
			t.Error(err)
		}
		if resDel.DeletedCount != 0 {
			t.Error(fmt.Errorf("deleted item count: " +
				strconv.FormatInt(resDel.DeletedCount, 10) + ", should be 0."))
		}
	})
}

func TestDeleteExpense(t *testing.T) {
	// db connection
	config.Init()
	db.Init()
	defer db.Close()

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
		if err != nil {
			t.Error(err)
		}
		if resDel.DeletedCount != 0 {
			t.Error(fmt.Errorf("deleted expense count: " +
				strconv.FormatInt(resDel.DeletedCount, 10) + ", should be 0."))
		}
	})

	t.Run("valid data", func(t *testing.T) {
		resDel, err := exp.DeleteExpense(expenseId)
		if err != nil {
			t.Errorf(err.Error())
		}
		if (*resDel).DeletedCount != 1 {
			t.Error(
				fmt.Errorf("deleted count should be 1, instead of %d",
					(*resDel).DeletedCount),
			)
		}
	})
}
