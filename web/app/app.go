package app

import (
	db "expense-logger/web/app/models"
	router "expense-logger/web/app/router"
)

func Init() {
	db.Init()
	router.Init()
}

func Run() {
	router.Run()
}
