package expense

import (
	"encoding/json"
	exp "expense-logger/web/app/models/expense"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func PutExpense(c *gin.Context) {
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

	var resItemInsert *mongo.InsertManyResult
	var resItemUpdate *mongo.InsertManyResult
	var resExpUpdate *mongo.UpdateResult

	// insert non-exist items
	if len(payload.Items) != 0 {
		var items []exp.Item
		var indexes []int
		for idx, it := range payload.Items {
			if it.Id.IsZero() {
				tmp := it.parse()
				tmp.ExpenseID = payload.Id
				items = append(items, tmp)
				indexes = append(indexes, idx)
			}
		}
		resItemInsert, err = exp.InsertItem(items)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		}
		// update id of newly inserted items
		for idx, id := range resItemInsert.InsertedIDs {
			index := indexes[idx]
			payload.Items[index].Id = id.(primitive.ObjectID)
		}
	}

	// update all items
	itemsUpdate := make([]exp.Item, 0, len(payload.Items))
	for _, it := range payload.Items {
		itemsUpdate = append(itemsUpdate, it.parse())
	}
	resItemUpdate, err = exp.UpdateItem(itemsUpdate)

	// update expense
	resExpUpdate, err = exp.UpdateExpense(payload.parse())
}
