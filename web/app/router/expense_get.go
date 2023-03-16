package router

import (
	"encoding/json"
	"io"
	"net/http"

	exp "expense-logger/web/app/models/expense"

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
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		err = json.Unmarshal(body, &param)
	} else {
		param.Count = 8
		param.Offset = 0
	}

	expenseSet, err := exp.GetExpenseLatest(uint8(param.Count), uint8(param.Offset))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": expenseSet,
	})
}
