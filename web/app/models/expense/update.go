package expense

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateExpense(newExpense Expense) (*mongo.UpdateResult, error) {
	filter := bson.D{
		{Key: "id", Value: newExpense.ExpenseID},
	}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			// something
		}},
	}
	res, err := expCollection.UpdateOne(
		ctx,
		filter,
		update,
	)

	if err != nil {
		return nil, err
	}
	return res, nil

}
