package models

import (
	"envelope/models/validators"
	"strconv"
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

func FromMap(data map[string][]string) (*Message, error) {
	err := validators.ValidateMessageMap(data)
	if err != nil {
		return nil, err
	}
	message := &Message{}
	message.Content = data["content"][0]
	unlocksAt, _ := strconv.Atoi(data["unlocks_at"][0])
	message.UnlocksAt = int64(unlocksAt)
	return message, nil
}
