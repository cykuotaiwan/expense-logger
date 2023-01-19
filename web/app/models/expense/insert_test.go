package expense_test

import (
	config "expense-logger/configs"
	db "expense-logger/web/app/models"
	exp "expense-logger/web/app/models/expense"
	"fmt"

	"testing"
)

func TestInsert(t *testing.T) {
	// db connection
	config.Init()
	db.Init()
	defer db.Close()

	// setup test data
	var newItems = []exp.Item{
		{
			ItemID: 1,
			Name:   "Potatoes",
			Price:  7.99,
			Count:  20,
		},
		{
			ItemID: 9887,
			Name:   "Milk",
			Price:  13.98,
			Count:  1,
		},
	}
	// insert items
	res, err := exp.InsertItem(newItems)
	fmt.Println(res)
	if err != nil {
		fmt.Println(err.Error())
		t.Errorf(err.Error())
	}

}
