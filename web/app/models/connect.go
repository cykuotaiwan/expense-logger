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

var client *mongo.Client
var ctx context.Context

func Connect() {
	username := os.Getenv("DBUSER")
	password := os.Getenv("DBPWD")
	domain := os.Getenv("DBDOMAIN")
	option := os.Getenv("DBOPTION")

	url := "mongodb+srv://" + username + ":" + password +
		"@" + domain + "/?" + option
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// list db
	dbs, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dbs)

	books := client.Database("test").Collection("books")
	cursor, err := books.Find(ctx, bson.D{})
	var result []bson.M
	if err = cursor.All(ctx, &result); err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(result))

}

func Disconnect() {
	client.Disconnect(ctx)
}
