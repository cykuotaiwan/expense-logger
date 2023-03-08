package router_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	rt "expense-logger/web/app/router"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetExpenseLatest(t *testing.T) {
	router := gin.Default()
	router.GET("/expense_latest", rt.GetExpenseLatest)

	writer := httptest.NewRecorder()

	t.Run("with payload", func(t *testing.T) {
		payload := map[string]interface{}{"cnt": 90, "offset": 0}
		jsonObj, _ := json.Marshal(payload)
		req, _ := http.NewRequest("GET", "/expense_latest", bytes.NewBuffer(jsonObj))

		router.ServeHTTP(writer, req)

		assert.Equal(t, http.StatusOK, writer.Code)
	})

	t.Run("without payload", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/expense_latest", nil)

		router.ServeHTTP(writer, req)

		assert.Equal(t, http.StatusOK, writer.Code)
	})

}
