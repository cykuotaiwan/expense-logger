package expense

import (
	"encoding/json"
	exp "expense-logger/web/app/models/expense"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostExpense(c *gin.Context) {
	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	var payload ExpensePayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	// insert items
	items := make([]exp.Item, len(payload.Items))
	for index, it := range payload.Items {
		items[index] = it.parse()
	}
	resItem, err := exp.InsertItem(items)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	if len(resItem.InsertedIDs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "item count shouldn't be zero",
		})
		return
	}

	// insert expense
	expense := payload.parseWithInsertResult(resItem)
	resExp, err := exp.InsertExpense(&expense)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	// success response
	insertedItemLen := len(resItem.InsertedIDs)
	insertedExpenseLen := 0
	if resExp != nil {
		insertedExpenseLen = 1
	}
	c.JSON(http.StatusOK, gin.H{
		"itemInsertedCnt":    insertedItemLen,
		"expenseInsertedCnt": insertedExpenseLen,
	})
}
