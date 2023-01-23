package expense

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateItem(newItem Item, itemId primitive.ObjectID) (*mongo.UpdateResult, error) {
	filter := bson.M{
		"_id": bson.M{"$eq": itemId},
	}

	update := bson.D{
		{
			Key: "$set",
			Value: bson.M{
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

func UpdateExpense(newExpense Expense, expId primitive.ObjectID) (*mongo.UpdateResult, error) {
	filter := bson.M{
		"_id": bson.M{"$eq": expId},
	}

	update := bson.D{
		{
			Key: "$set",
			Value: bson.M{
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
