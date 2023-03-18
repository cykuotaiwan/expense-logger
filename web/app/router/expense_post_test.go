package router_test

import (
	"bytes"
	"encoding/json"
	exp "expense-logger/web/app/models/expense"
	rt "expense-logger/web/app/router"
	util "expense-logger/web/app/util"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPostExpense(t *testing.T) {
	util.NewDBConnection()
	defer util.EndDBConnection()

	path := "/item_latest"
	router := gin.Default()
	router.POST(path, rt.PostExpense)

	var items = []rt.ItemPayload{
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
	var expense = rt.ExpensePayload{
		DateTime: time.Now(),
		ShopName: "Loblaws",
		Total:    2011,
		Items:    items,
	}

	t.Run("with valid payload", func(t *testing.T) {
		payload, _ := json.Marshal(expense)

		writer, err := util.PerformRequest(router, "POST", path, bytes.NewBuffer(payload))
		assert.Equal(t, http.StatusOK, writer.Code)
		assert.Nil(t, err)

		var rsp map[string]int
		err = json.Unmarshal([]byte(writer.Body.Bytes()), &rsp)
		itemCnt, itemExist := rsp["itemInsertedCnt"]
		expCnt, expExist := rsp["expenseInsertedCnt"]

		assert.Nil(t, err)
		assert.Greater(t, 0, itemCnt)
		assert.Greater(t, 0, expCnt)
		assert.True(t, itemExist)
		assert.True(t, expExist)
	})
}
