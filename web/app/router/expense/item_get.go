package expense

import (
	"encoding/json"
	exp "expense-logger/web/app/models/expense"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetItemLatest(c *gin.Context) {
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

	itemSet, err := exp.GetItemLatest(uint8(param.Count), uint8(param.Offset))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": itemSet,
	})
}
