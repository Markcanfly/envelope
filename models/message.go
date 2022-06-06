package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Message struct {
	ID 	primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Content string `json:"content" bson:"content"`
	UnlocksAt int64 `json:"unlocksAt" bson:"unlocksAt"`
	//User *User `json:"_id,omitempty" bson:"_id,omitempty"`
}
