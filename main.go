package main

import (
	config "expense-logger/configs"
	app "expense-logger/web/app"
)

func init() {
	config.Init()
	app.Init()
}

func main() {
	app.Run()
}
