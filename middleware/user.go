package middleware

import (
	"context"
	"envelope/models"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	err := createUser(username, email, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func createUser(username, email, password string) error {
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

