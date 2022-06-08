package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Message struct {
	ID 	primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Content string `json:"content" bson:"content"`
	UnlocksAt int64 `json:"unlocksAt,omitempty" bson:"unlocksAt,omitempty"`
	//User *User `json:"_id,omitempty" bson:"_id,omitempty"`
}

func (m *Message) IsOpened() bool {
	return m.UnlocksAt != 0 && m.UnlocksAt < time.Now().Unix()
}
