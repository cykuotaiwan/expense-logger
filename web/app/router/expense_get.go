package router

import (
	"encoding/json"
	exp "expense-logger/web/app/models/expense"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type latestParam struct {
	Count  int `json:"cnt"`
	Offset int `json:"offset"`
}

type getByDateParam struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

func GetExpenseLatest(c *gin.Context) {
	var param latestParam

	if c.Request.Body != nil {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		err = json.Unmarshal(body, &param)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		if param.Offset < 0 || param.Count < 0 {
			c.JSON(http.StatusBadRequest, nil)
			return
		}
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

func GetExpenseByDate(c *gin.Context) {
	var param getByDateParam

	if c.Request.Body != nil {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		err = json.Unmarshal(body, &param)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
	} else {
		// one day
		param.StartDate = time.Now().AddDate(0, 0, -1)
		param.EndDate = time.Now()
	}

	expenseSet, err := exp.GetExpenseByDate(param.StartDate, param.EndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": expenseSet,
	})
}
