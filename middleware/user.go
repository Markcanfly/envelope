package middleware

import (
	"context"
	"envelope/models"
	"errors"
	"log"
)

func CreateUser(username, email, password string) error {
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

// Get token for user
func AuthenticateUser(username, password string) (string, error) {
	var user *models.User
	err := userCollection.FindOne(context.Background(), &models.User{Username: username}).Decode(&user)
	if err != nil {
		return "", errors.New("invalid username or password")
	}
	err = user.CheckPassword(password)
	if err != nil {
		return "", errors.New("invalid username or password")
	}
	token := models.CreateTokenFor(user)
	return token, nil
}
