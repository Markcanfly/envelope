package middleware

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var messageCollection *mongo.Collection
var userCollection *mongo.Collection

func InitDb(testing bool) {
	connectToDb(testing)
}

func connectToDb(testing bool) {
	var uri string
	if testing {
		uri = os.Getenv("TEST_MONGODB_URI")
	} else {
		uri = os.Getenv("MONGODB_URI")
	}
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI'/'TEST_MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("envelope")
	messageCollection = db.Collection("messages")
	userCollection = db.Collection("users")
}

func DisconnectDb() {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
