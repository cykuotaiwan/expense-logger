package expense

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateItem(newItem Item) (*mongo.UpdateResult, error) {
	filter := bson.M{
		"_id": bson.M{"$eq": newItem.Id},
	}

	update := bson.D{
		{
			Key: "$set",
			Value: bson.M{
				"_id":       newItem.Id,
				"name":      newItem.Name,
				"price":     newItem.Price,
				"count":     newItem.Count,
				"unit":      newItem.Unit,
				"expenseId": newItem.ExpenseID,
			},
		},
	}

	res, err := itemCollection.UpdateOne(
		context.TODO(),
		filter,
		update,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func UpdateExpense(newExpense Expense) (*mongo.UpdateResult, error) {
	filter := bson.M{
		"_id": bson.M{"$eq": newExpense.Id},
	}

	update := bson.D{
		{
			Key: "$set",
			Value: bson.M{
				"_id":      newExpense.Id,
				"datetime": newExpense.DateTime,
				"shopName": newExpense.ShopName,
				"total":    newExpense.Total,
				"itemIDs":  newExpense.ItemIDs,
			},
		},
	}
	res, err := expCollection.UpdateOne(
		context.TODO(),
		filter,
		update,
	)

	if err != nil {
		return nil, err
	}
	return res, nil
}
