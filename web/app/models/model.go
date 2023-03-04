package models

import (
	expense "expense-logger/web/app/models/expense"
	user "expense-logger/web/app/models/user"

	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var ctx context.Context

func Connect() {
	var url string

	mode := os.Getenv("DBCONNMODE")
	if mode == "LOCAL" {
		port := os.Getenv("DBLOCALPORT")
		url = "mongodb://localhost:" + port
	} else if mode == "REMOTE" {
		username := os.Getenv("DBUSER")
		password := os.Getenv("DBPWD")
		domain := os.Getenv("DBDOMAIN")
		option := os.Getenv("DBOPTION")

		url = "mongodb+srv://" + username + ":" + password +
			"@" + domain + "/?" + option
	} else {
		log.Fatal("unsupported connect mode: " + mode)
	}

	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func Disconnect() {
	client.Disconnect(ctx)
}

func Init() {
	Connect()
	expense.SetCollections(client, &ctx)
	user.SetCollections(client, &ctx)
}

func Close() {
	Disconnect()
}
