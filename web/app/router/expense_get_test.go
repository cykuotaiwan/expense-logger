package router_test

import (
	"bytes"
	"encoding/json"
	exp "expense-logger/web/app/models/expense"
	rt "expense-logger/web/app/router"
	util "expense-logger/web/app/util"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetExpenseLatest(t *testing.T) {
	util.NewDBConnection()
	defer util.EndDBConnection()

	path := "/expense_latest"
	router := gin.Default()
	router.GET(path, rt.GetExpenseLatest)

	t.Run("with payload", func(t *testing.T) {
		param := map[string]interface{}{"cnt": 15, "offset": 0}
		payload, _ := json.Marshal(param)

		// req, _ := http.NewRequest("GET", "/expense_latest", bytes.NewBuffer(jsonObj))
		writer, err := util.PerformRequest(router, "GET", path, bytes.NewBuffer(payload))

		assert.Equal(t, http.StatusOK, writer.Code)
		assert.Nil(t, err)

		var rsp map[string][]exp.Expense
		err = json.Unmarshal([]byte(writer.Body.String()), &rsp)
		expArr, exist := rsp["data"]

		assert.Nil(t, err)
		assert.True(t, exist)
		assert.Equal(t, len(expArr), 15)
	})

	t.Run("without payload", func(t *testing.T) {
		writer, err := util.PerformRequest(router, "GET", path, nil)

		assert.Equal(t, http.StatusOK, writer.Code)
		assert.Nil(t, err)

		var rsp map[string][]exp.Expense
		err = json.Unmarshal([]byte(writer.Body.String()), &rsp)
		expArr, exist := rsp["data"]

		assert.Nil(t, err)
		assert.True(t, exist)
		assert.Equal(t, len(expArr), 8)
	})

}
