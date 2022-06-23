package middleware

import (
	"context"
	"envelope/models"
	"errors"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	r.ParseForm()
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	err := createUser(username, email, password)
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		return
	}
	switch err.Error() {
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	case "email is already registered":
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	case "username is already taken":
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func createUser(username, email, password string) error {
	if CheckEmailRegistered(email) {
		return errors.New("email is already registered") // TODO error class
	}
	if CheckUsernameTaken(username) {
		return errors.New("username is already taken") // TODO error class
	}
	user, err := models.NewUser(username, email, password)
	if err != nil {
		return err
	}
	_, err = userCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// TODO CSRF
// Check if username is already taken in our system
func CheckUsernameTaken(username string) bool {
	result := userCollection.FindOne(context.Background(), bson.D{primitive.E{Key: "username", Value: username}})
	if result.Err() == mongo.ErrNoDocuments {
		return false
	} else if result.Err() == nil {
		return true
	} else {
		log.Fatal("Unexpected error trying to fetch username" + username + ": " + result.Err().Error())
		return true
	}
}

// TODO CSRF
// Check if email is already registered in our system
func CheckEmailRegistered(email string) bool {
	result := userCollection.FindOne(context.Background(), bson.D{primitive.E{Key: "email", Value: email}})
	if result.Err() == mongo.ErrNoDocuments {
		return false
	} else if result.Err() == nil {
		return true
	} else {
		log.Fatal("Unexpected error trying to fetch email" + email + ": " + result.Err().Error())
		return false
	}
}
