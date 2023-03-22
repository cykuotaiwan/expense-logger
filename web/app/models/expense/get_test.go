package expense_test

import (
	exp "expense-logger/web/app/models/expense"
	util "expense-logger/web/app/util"
	"time"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetExpenseLatest(t *testing.T) {
	util.NewDBConnection()
	defer util.EndDBConnection()

	expense, err := exp.GetExpenseLatest(30, 0)
	assert.Nil(t, err)
	assert.NotNil(t, expense)
}

func TestGetItemLatest(t *testing.T) {
	util.NewDBConnection()
	defer util.EndDBConnection()

	t.Run("query zero", func(t *testing.T) {
		items, err := exp.GetItemLatest(0, 0)

		assert.Nil(t, err)
		assert.Nil(t, items)
	})
	t.Run("valid value", func(t *testing.T) {
		var queryCnt uint8 = 30
		items, err := exp.GetItemLatest(queryCnt, 0)

		assert.Nil(t, err)
		assert.NotNil(t, items)
		assert.LessOrEqual(t, 30, len(items))
	})

}

func TestGetExpenseByDate(t *testing.T) {
	util.NewDBConnection()
	defer util.EndDBConnection()

	t.Run("late start time", func(t *testing.T) {
		end := time.Now()
		start := time.Now().AddDate(0, 0, 1)
		expense, err := exp.GetExpenseByDate(start, end)

		assert.NotNil(t, err)
		assert.Nil(t, expense)
	})

	t.Run("start equal end", func(t *testing.T) {
		end := time.Now()
		start := end
		expense, err := exp.GetExpenseByDate(start, end)

		assert.NotNil(t, err)
		assert.Nil(t, expense)
	})

	t.Run("valid value", func(t *testing.T) {
		start := time.Date(2023, time.January, 21, 0, 0, 0, 0, time.UTC)
		end := start.AddDate(0, 0, 1)
		_, err := exp.GetExpenseByDate(start, end)

		assert.Nil(t, err)
	})
}
