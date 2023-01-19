package expense

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func InsertItem(newItems []Item) (*mongo.InsertManyResult, error) {
	if (newItems == nil) || (len(newItems) == 0) {
		return nil, nil
	}

	items := make([]interface{}, len(newItems))
	size := 0
	for i := range newItems {
		if !newItems[i].IsEmpty() {
			items[i] = newItems[i]
			size++
		}
	}
	if size == 0 {
		return nil, nil
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
	if (newExpense == nil) || (*newExpense).IsEmpty() {
		return nil, nil
	}

	res, err := expCollection.InsertOne(
		context.TODO(),
		newExpense,
	)

	if err != nil {
		return nil, err
	}
	return res, nil
}
