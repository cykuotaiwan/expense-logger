package expense

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateExpense(newExpense Expense, expId primitive.ObjectID) (*mongo.UpdateResult, error) {
	filter := bson.D{
		{Key: "id", Value: expId},
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
