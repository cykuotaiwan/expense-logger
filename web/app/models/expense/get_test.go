package expense_test

import (
	config "expense-logger/configs"
	db "expense-logger/web/app/models"
	exp "expense-logger/web/app/models/expense"

	"testing"
)

func TestGetExpense(t *testing.T) {
	// db connection
	config.Init()
	db.Init()
	defer db.Close()

	expense, err := exp.GetExpenseLatest(30, 0)
	if err != nil || expense == nil {
		t.Errorf(err.Error())
	}
}

func TestGetItem(t *testing.T) {
	// db connection
	config.Init()
	db.Init()
	defer db.Close()

	items, err := exp.GetItemLatest(30, 0)
	if err != nil || items == nil {
		t.Errorf(err.Error())
	}
}
