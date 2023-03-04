package user

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserGoogle struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	ProfilePic   string `json:"profilePic"`
	ExternalType string `json:"externalType"`
	ExternalID   string `json:"externalID"`
}

var userCollection *mongo.Collection
var ctx context.Context

func SetCollections(client *mongo.Client, context *context.Context) {
	if os.Getenv("OPMODEs") == "RELEASE" {
		userCollection = client.Database("user").Collection("user")
	} else {
		userCollection = client.Database("user-debug").Collection("user")
	}

	ctx = *context
}

func (user UserGoogle) IsEmpty() bool {
	empty := false
	empty = empty || len(user.ID) == 0
	empty = empty || len(user.Username) == 0
	empty = empty || len(user.Email) == 0
	empty = empty || len(user.ProfilePic) == 0
	empty = empty || len(user.ExternalType) == 0
	empty = empty || len(user.ExternalID) == 0

	return empty
}
