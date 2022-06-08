package models

import (
	"envelope/models/validators"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Email string `json:"email" bson:"email"`
	PasswordHash string `json:"passwordhash" bson:"passwordhash"`
}

func (u *User) SetPassword(password string) error {
	if err := validators.ValidatePassword(password); err != nil {
		return err
	}
	passwordHash, err := GenerateHashFromPassword(password)
	if err != nil {
		u.PasswordHash = string(passwordHash)
	}
	return nil
}

func GenerateHashFromPassword(password string) ([]byte, error) {
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	return passwordHash, err
}

func (u *User) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func NewUser(username, email, password string) (*User, error) {
	if err := validators.ValidateUsername(username); err != nil { return nil, err }
	if err := validators.ValidatePassword(password); err != nil { return nil, err }
	if err := validators.ValidateEmail(email);       err != nil { return nil, err }
	passwordHash, err := GenerateHashFromPassword(password)
	if err != nil { return nil, err }
	user := &User{
		Username: username,
		Email: email,
		PasswordHash: string(passwordHash),
	}
	return user, nil
}
