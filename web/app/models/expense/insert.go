package expense

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func InsertItem(newItems []Item) (*mongo.InsertManyResult, error) {
	items := make([]interface{}, len(newItems))
	for i := range newItems {
		items[i] = newItems[i]
	}
	res, err := itemCollection.InsertMany(
		context.TODO(),
		items,
	)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func InsertExpense(newExpense *Expense) (*mongo.InsertOneResult, error) {
	res, err := expCollection.InsertOne(
		ctx,
		newExpense,
	)

	if err != nil {
		return nil, err
	}
	return res, nil
}
