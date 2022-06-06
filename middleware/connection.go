package middleware

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO swap out for separate functions for db setup, and a singleton Collection var

var client *mongo.Client
var messageCollection *mongo.Collection
//var userCollection *mongo.Collection

func InitDb() {
	loadEnv()
	connectToDb()
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func connectToDb() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
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
	//userCollection = db.Collection("users")
}

func DisconnectDb() {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
