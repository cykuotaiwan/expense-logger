package expense

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteItem(itemId primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": itemId}

	res, err := itemCollection.DeleteOne(
		context.TODO(),
		filter,
	)
	if err != nil {
		return nil, err
	}

	if (*res).DeletedCount != 0 {
		cursor, err := expCollection.Find(
			context.TODO(),
			bson.M{
				"itemIDs": itemId,
			},
		)
		if err == nil {
			var expenses []Expense
			if err = cursor.All(context.TODO(), &expenses); err == nil {
				for _, exp := range expenses {
					var tmp []primitive.ObjectID
					for _, id := range exp.ItemIDs {
						if id != itemId {
							tmp = append(tmp, id)
						}
						exp.ItemIDs = tmp
					}
					expCollection.UpdateOne(
						context.TODO(),
						bson.M{"_id": bson.M{"$eq": exp.Id}},
						bson.D{
							{
								Key: "$set",
								Value: bson.M{
									"_id":      exp.Id,
									"datetime": exp.DateTime,
									"shopName": exp.ShopName,
									"total":    exp.Total,
									"itemIDs":  exp.ItemIDs,
								},
							},
						},
					)
				}
			}
		}
	}
	return res, nil
}

func DeleteExpense(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}

	res, err := itemCollection.DeleteOne(
		context.TODO(),
		filter,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
