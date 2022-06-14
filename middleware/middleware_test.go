package middleware

import (
	"testing"
)

func cleanTestDatabase() {
	CleanDatabase()
}

var token string
func TestMain(m *testing.M) {
	InitDb(true)
	email, user, pw := "test@example.com", "test", "bare_minimum_password"
	helpCreateUser(email, user, pw)
	rr := helpLogin(email, pw)
	token = rr.Header()["Authorization"][0]
	if token == "" {
		panic("token is empty")
	}
	m.Run()
	cleanTestDatabase()
}
