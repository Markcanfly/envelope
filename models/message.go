package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Message struct {
	ID 	primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Content string `json:"content" bson:"content"`
	UnlocksAt int64 `json:"unlocks_at,omitempty" bson:"unlocks_at,omitempty"` // should be *int64 so it can be nil
	User primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
}

func (m *Message) IsOpened() bool {
	return m.UnlocksAt != 0 && m.UnlocksAt < time.Now().Unix()
}
