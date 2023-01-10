package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect() {
	username := os.Getenv("DBUSER")
	password := os.Getenv("DBPWD")
	domain := os.Getenv("DBDOMAIN")
	option := os.Getenv("DBOPTION")

	url := "mongodb+srv://" + username + ":" + password + "@" + domain + "/?" + option
	Client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer Client.Disconnect(ctx)

	// list db
	dbs, err := Client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dbs)

	books := Client.Database("test").Collection("books")
	cursor, err := books.Find(ctx, bson.D{})
	var result []bson.M
	if err = cursor.All(ctx, &result); err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(result))

	// Client.Disconnect(ctx)
}
