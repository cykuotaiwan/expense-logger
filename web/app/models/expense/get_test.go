package expense_test

import (
	config "expense-logger/configs"
	db "expense-logger/web/app/models"
	exp "expense-logger/web/app/models/expense"
	"fmt"
	"time"

	"testing"
)

func TestGetExpenseLatest(t *testing.T) {
	// db connection
	config.Init()
	db.Init()
	defer db.Close()

	expense, err := exp.GetExpenseLatest(30, 0)
	if err != nil || expense == nil {
		t.Errorf(err.Error())
	}
}

func TestGetItemLatest(t *testing.T) {
	// db connection
	config.Init()
	db.Init()
	defer db.Close()

	t.Run("query zero", func(t *testing.T) {
		items, err := exp.GetItemLatest(0, 0)
		if err != nil || items != nil {
			t.Errorf(err.Error())
		}
	})
	t.Run("valid value", func(t *testing.T) {
		var queryCnt uint8 = 30
		items, err := exp.GetItemLatest(queryCnt, 0)
		if err != nil || items == nil {
			t.Errorf(err.Error())
		}
	})

}

func TestGetExpenseByDate(t *testing.T) {
	// db connection
	config.Init()
	db.Init()
	defer db.Close()

	t.Run("late start time", func(t *testing.T) {
		end := time.Now()
		start := time.Now().AddDate(0, 0, 1)
		expense, err := exp.GetExpenseByDate(start, end)
		if err == nil {
			t.Error(fmt.Errorf("should have an error"))
		}
		if expense != nil {
			t.Error(fmt.Errorf("expense should return nil"))
		}
	})

	t.Run("start equal end", func(t *testing.T) {
		end := time.Now()
		start := end
		expense, err := exp.GetExpenseByDate(start, end)
		if err == nil {
			t.Error(fmt.Errorf("should have an error"))
		}
		if expense != nil {
			t.Error(fmt.Errorf("expense should return nil"))
		}
	})

	t.Run("valid value", func(t *testing.T) {
		start := time.Date(2023, time.January, 21, 0, 0, 0, 0, time.UTC)
		end := start.AddDate(0, 0, 1)
		expense, err := exp.GetExpenseByDate(start, end)
		if err != nil || len(expense) == 0 {
			t.Errorf(err.Error())
		}
		if len(expense) == 0 {
			t.Error(fmt.Errorf("expense lenght should not be empty"))
		}
	})
}
