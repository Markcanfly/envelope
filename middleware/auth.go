package middleware

import (
	"context"
	"envelope/models"
	"errors"
	"net/http"
)

const authErrorMessage = "invalid username or password"

// TODO XSS protection
func Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	token, err := authenticateUser(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Authorization", token)
}

func TokenAuth(w http.ResponseWriter, r *http.Request) (*models.TokenData, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return nil, errors.New("unauthorized")
	}
	tokenData := models.GetTokenData(token)
	if tokenData == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return nil, errors.New("unauthorized")
	}
	return tokenData, nil
}

// Get token for user
func authenticateUser(username, password string) (string, error) {
	var user *models.User
	err := userCollection.FindOne(context.Background(), &models.User{Username: username}).Decode(&user)
	if err != nil {
		return "", errors.New(authErrorMessage)
	}
	err = user.CheckPassword(password)
	if err != nil {
		return "", errors.New(authErrorMessage)
	}
	token := models.CreateTokenFor(user)
	return token, nil
}

