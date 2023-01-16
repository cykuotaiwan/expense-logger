package expense

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetExpenseLatest(cnt int8, offset int8) ([]expense, error) {
	option := options.Find()
	option.SetSort(bson.M{"$natural": -1})
	option.SetLimit(int64(cnt))
	filter := bson.D{}
	cur, err := expCollection.Find(ctx, filter, option)
	if err != nil {
		return nil, err
	}
	var res []expense
	if err = cur.All(ctx, &res); err != nil {
		return nil, err
	}
	return res[offset:], nil
}
