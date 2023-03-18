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

func TestGetItemLatest(t *testing.T) {
	util.NewDBConnection()
	defer util.EndDBConnection()

	path := "/item_latest"
	router := gin.Default()
	router.GET(path, rt.GetItemLatest)

	t.Run("with valid payload", func(t *testing.T) {
		param := map[string]interface{}{"cnt": 15, "offset": 0}
		payload, _ := json.Marshal(param)

		writer, err := util.PerformRequest(router, "GET", path, bytes.NewBuffer(payload))

		assert.Equal(t, http.StatusOK, writer.Code)
		assert.Nil(t, err)

		var rsp map[string][]exp.Item
		err = json.Unmarshal([]byte(writer.Body.Bytes()), &rsp)
		itemArr, exist := rsp["data"]

		assert.Nil(t, err)
		assert.True(t, exist)
		assert.LessOrEqual(t, 15, len(itemArr))
	})

	t.Run("without payload", func(t *testing.T) {
		writer, err := util.PerformRequest(router, "GET", path, nil)

		assert.Equal(t, http.StatusOK, writer.Code)
		assert.Nil(t, err)

		var rsp map[string][]exp.Item
		err = json.Unmarshal([]byte(writer.Body.Bytes()), &rsp)
		itemArr, exist := rsp["data"]

		assert.Nil(t, err)
		assert.True(t, exist)
		assert.LessOrEqual(t, 8, len(itemArr))
	})
	t.Run("with valid payload 0 count", func(t *testing.T) {
		param := map[string]interface{}{"cnt": 0, "offset": 0}
		payload, _ := json.Marshal(param)

		writer, err := util.PerformRequest(router, "GET", path, bytes.NewBuffer(payload))

		assert.Equal(t, http.StatusOK, writer.Code)
		assert.Nil(t, err)

		var rsp map[string][]exp.Item
		err = json.Unmarshal([]byte(writer.Body.Bytes()), &rsp)
		itemArr, exist := rsp["data"]

		assert.Nil(t, err)
		assert.True(t, exist)
		assert.Equal(t, 0, len(itemArr))
	})
	t.Run("with invalid payload negetive count", func(t *testing.T) {
		param := map[string]interface{}{"cnt": -1, "offset": 0}
		payload, _ := json.Marshal(param)

		writer, err := util.PerformRequest(router, "GET", path, bytes.NewBuffer(payload))

		assert.Equal(t, http.StatusBadRequest, writer.Code)
		assert.Nil(t, err)

		var rsp map[string][]exp.Item
		err = json.Unmarshal([]byte(writer.Body.Bytes()), &rsp)
		itemArr, exist := rsp["data"]

		assert.Nil(t, err)
		assert.False(t, exist)
		assert.Equal(t, 0, len(itemArr))
	})
	t.Run("with invalid payload negetive offset", func(t *testing.T) {
		param := map[string]interface{}{"cnt": 0, "offset": -1}
		payload, _ := json.Marshal(param)

		writer, err := util.PerformRequest(router, "GET", path, bytes.NewBuffer(payload))

		assert.Equal(t, http.StatusBadRequest, writer.Code)
		assert.Nil(t, err)

		var rsp map[string][]exp.Item
		err = json.Unmarshal([]byte(writer.Body.Bytes()), &rsp)
		itemArr, exist := rsp["data"]

		assert.Nil(t, err)
		assert.False(t, exist)
		assert.Equal(t, 0, len(itemArr))
	})
}
