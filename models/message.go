package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Message struct {
	ID 	primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Content string `json:"content,omitempty"`
	UnlocksAt int64 `json:"unlocksAt,omitempty" bson:"unlocksAt,omitempty"`
	//User *User `json:"_id,omitempty" bson:"_id,omitempty"`
}
