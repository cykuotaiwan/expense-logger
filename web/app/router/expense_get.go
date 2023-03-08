package router

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type expenseParam struct {
	Count  int `json:"cnt"`
	Offset int `json:"offset"`
}

func GetExpenseLatest(c *gin.Context) {
	var param expenseParam

	if c.Request.Body != nil {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, nil)
		}
		err = json.Unmarshal(body, &param)
	} else {
		param.Count = 8
		param.Offset = 0
	}
	c.IndentedJSON(http.StatusOK, nil)
}
