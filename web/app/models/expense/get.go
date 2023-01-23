package expense

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetExpenseByDate(startDate, endDate time.Time) ([]Expense, error) {
	if startDate.After(endDate) {
		return nil, fmt.Errorf("Invalid Parameters: startDate should be earlier than endDate")
	}
	if startDate.Equal(endDate) {
		return nil, fmt.Errorf("Invalid Parameters: startDate should be earlier than endDate")
	}

	option := options.Find()
	option.SetSort(bson.M{"$natural": -1})
	filter := bson.M{
		"datetime": bson.M{
			"$gte": primitive.NewDateTimeFromTime(startDate),
			"$lt":  primitive.NewDateTimeFromTime(endDate),
		},
	}

	cursor, err := expCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	var res []Expense
	if err = cursor.All(ctx, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func GetItemLatest(cnt uint8, offset uint8) ([]Item, error) {
	if cnt == 0 {
		return nil, nil
	}

	option := options.Find()
	option.SetSort(bson.M{"$natural": -1})
	option.SetLimit(int64(cnt + offset))
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
