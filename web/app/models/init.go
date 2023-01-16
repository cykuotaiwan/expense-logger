package models

import (
	expense "expense-logger/web/app/models/expense"
)

func Init() {
	Connect()
	expense.SetExpenseCollection(client, &ctx)
}
