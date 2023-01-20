package expense

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetExpenseLatest(cnt uint8, offset uint8) ([]Expense, error) {
	option := options.Find()
	option.SetSort(bson.M{"$natural": -1})
	option.SetLimit(int64(cnt))
	filter := bson.D{}

	cursor, err := expCollection.Find(context.TODO(), filter, option)
	if err != nil {
		return nil, err
	}

	var res []Expense
	if err = cursor.All(ctx, &res); err != nil {
		return nil, err
	}

	return res[offset:], nil
}

func GetItemLatest(cnt uint8, offset uint8) ([]Item, error) {
	option := options.Find()
	option.SetSort(bson.M{"$natural": -1})
	option.SetLimit(int64(cnt))
	filter := bson.D{}

	cursor, err := itemCollection.Find(context.TODO(), filter, option)
	if err != nil {
		return nil, err
	}

	var res []Item
	if err = cursor.All(ctx, &res); err != nil {
		return nil, err
	}
	return res[offset:], nil
}
