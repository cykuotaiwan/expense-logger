package expense_test

import (
	"bytes"
	"encoding/json"
	exp "expense-logger/web/app/models/expense"
	rt "expense-logger/web/app/router/expense"
	util "expense-logger/web/app/util"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetExpenseLatest(t *testing.T) {
	util.NewDBConnection()
	defer util.EndDBConnection()

	path := "/expense_latest"
	router := gin.Default()
	router.GET(path, rt.GetExpenseLatest)

	t.Run("with valid payload", func(t *testing.T) {
		param := map[string]interface{}{"cnt": 15, "offset": 0}
		payload, _ := json.Marshal(param)

		writer, err := util.PerformRequest(router, "GET", path, bytes.NewBuffer(payload))

		assert.Equal(t, http.StatusOK, writer.Code)
		assert.Nil(t, err)

		var rsp map[string][]exp.Expense
		err = json.Unmarshal([]byte(writer.Body.Bytes()), &rsp)
		expArr, exist := rsp["data"]

		assert.Nil(t, err)
		assert.True(t, exist)
		assert.LessOrEqual(t, 15, len(expArr))
	})
	t.Run("without payload", func(t *testing.T) {
		writer, err := util.PerformRequest(router, "GET", path, nil)

		assert.Equal(t, http.StatusOK, writer.Code)
		assert.Nil(t, err)

		var rsp map[string][]exp.Expense
		err = json.Unmarshal([]byte(writer.Body.Bytes()), &rsp)
		expArr, exist := rsp["data"]

		assert.Nil(t, err)
		assert.True(t, exist)
		assert.LessOrEqual(t, 8, len(expArr))
	})
	t.Run("with valid payload 0 count", func(t *testing.T) {
		param := map[string]interface{}{"cnt": 0, "offset": 0}
		payload, _ := json.Marshal(param)

		writer, err := util.PerformRequest(router, "GET", path, bytes.NewBuffer(payload))

		assert.Equal(t, http.StatusOK, writer.Code)
		assert.Nil(t, err)

		var rsp map[string][]exp.Expense
		err = json.Unmarshal([]byte(writer.Body.Bytes()), &rsp)
		expArr, exist := rsp["data"]

		assert.Nil(t, err)
		assert.True(t, exist)
		assert.Equal(t, 0, len(expArr))
	})
	t.Run("with invalid payload negetive count", func(t *testing.T) {
		param := map[string]interface{}{"cnt": -1, "offset": 0}
		payload, _ := json.Marshal(param)

		writer, err := util.PerformRequest(router, "GET", path, bytes.NewBuffer(payload))

		assert.Equal(t, http.StatusBadRequest, writer.Code)
		assert.Nil(t, err)

		var rsp map[string][]exp.Expense
		err = json.Unmarshal([]byte(writer.Body.Bytes()), &rsp)
		expArr, exist := rsp["data"]

		assert.Nil(t, err)
		assert.False(t, exist)
		assert.Equal(t, 0, len(expArr))
	})
	t.Run("with invalid payload negetive offset", func(t *testing.T) {
		param := map[string]interface{}{"cnt": 0, "offset": -1}
		payload, _ := json.Marshal(param)

		writer, err := util.PerformRequest(router, "GET", path, bytes.NewBuffer(payload))

		assert.Equal(t, http.StatusBadRequest, writer.Code)
		assert.Nil(t, err)

		var rsp map[string][]exp.Expense
		err = json.Unmarshal([]byte(writer.Body.Bytes()), &rsp)
		expArr, exist := rsp["data"]

		assert.Nil(t, err)
		assert.False(t, exist)
		assert.Equal(t, 0, len(expArr))
	})
}

func TestGetExpenseByDate(t *testing.T) {
	util.NewDBConnection()
	defer util.EndDBConnection()

	path := "/expense_date"
	router := gin.Default()
	router.GET(path, rt.GetExpenseByDate)

	t.Run("with valid payload", func(t *testing.T) {
		loc, _ := time.LoadLocation("UTC")
		param := map[string]interface{}{
			"startDate": time.Date(2023, 3, 3, 0, 0, 0, 0, loc),
			"endDate":   time.Date(2023, 3, 4, 0, 0, 0, 0, loc)}
		payload, _ := json.Marshal(param)

		writer, err := util.PerformRequest(router, "GET", path, bytes.NewBuffer(payload))
		assert.Equal(t, http.StatusOK, writer.Code)
		assert.Nil(t, err)

		var rsp map[string][]exp.Expense
		err = json.Unmarshal([]byte(writer.Body.Bytes()), &rsp)
		expArr, exist := rsp["data"]

		assert.Nil(t, err)
		assert.True(t, exist)
		assert.Equal(t, 6, len(expArr))
	})
	t.Run("without payload", func(t *testing.T) {
		writer, err := util.PerformRequest(router, "GET", path, nil)
		assert.Equal(t, http.StatusOK, writer.Code)
		assert.Nil(t, err)

		var rsp map[string][]exp.Expense
		err = json.Unmarshal([]byte(writer.Body.Bytes()), &rsp)
		expArr, exist := rsp["data"]

		assert.Nil(t, err)
		assert.True(t, exist)
		assert.Equal(t, 0, len(expArr))
	})
	t.Run("with invalid payload early end date", func(t *testing.T) {
		loc, _ := time.LoadLocation("UTC")
		param := map[string]interface{}{
			"startDate": time.Date(2023, 3, 4, 0, 0, 0, 0, loc),
			"endDate":   time.Date(2023, 3, 3, 0, 0, 0, 0, loc)}
		payload, _ := json.Marshal(param)

		writer, err := util.PerformRequest(router, "GET", path, bytes.NewBuffer(payload))
		assert.Equal(t, http.StatusInternalServerError, writer.Code)
		assert.Nil(t, err)

		var rsp map[string][]exp.Expense
		err = json.Unmarshal([]byte(writer.Body.Bytes()), &rsp)
		expArr, exist := rsp["data"]

		assert.Nil(t, err)
		assert.False(t, exist)
		assert.Equal(t, 0, len(expArr))
	})
	t.Run("with invalid payload same end date", func(t *testing.T) {
		loc, _ := time.LoadLocation("UTC")
		param := map[string]interface{}{
			"startDate": time.Date(2023, 3, 3, 0, 0, 0, 0, loc),
			"endDate":   time.Date(2023, 3, 3, 0, 0, 0, 0, loc)}
		payload, _ := json.Marshal(param)

		writer, err := util.PerformRequest(router, "GET", path, bytes.NewBuffer(payload))
		assert.Equal(t, http.StatusInternalServerError, writer.Code)
		assert.Nil(t, err)

		var rsp map[string][]exp.Expense
		err = json.Unmarshal([]byte(writer.Body.Bytes()), &rsp)
		expArr, exist := rsp["data"]

		assert.Nil(t, err)
		assert.False(t, exist)
		assert.Equal(t, 0, len(expArr))
	})
}
