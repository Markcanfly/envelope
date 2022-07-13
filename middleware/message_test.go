package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"envelope/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func helpCreateMessage(content string, timestamp int64) *httptest.ResponseRecorder {
	data := url.Values{}
	data.Set("content", content)
	data.Set("unlocks_at", fmt.Sprintf("%d", timestamp))

	req, _ := http.NewRequest("POST", "", bytes.NewBufferString(data.Encode()))
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateMessage)

	handler.ServeHTTP(rr, req)
	return rr
}

func TestTryCreateMessageWithoutAuth(t *testing.T) {
	message := `{"message":"titkos üzenet"}`
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(message))	
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateMessage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}
}

func TestCreateMessage(t *testing.T) {
	data := url.Values{}
	data.Set("content", "titkos üzenet")
	data.Set("unlocks_at", fmt.Sprintf("%d", time.Now().Unix()))
	
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(data.Encode()))	
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateMessage)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusCreated)
	}
}

func TestCreateMessageCheckData(t *testing.T) {
	currenttime := time.Now().Unix()
	rr := helpCreateMessage("titkos üzenet", currenttime)
	var message models.Message
	json.Unmarshal(rr.Body.Bytes(), &message)
	if message.Content != "titkos üzenet" {
		t.Errorf("handler returned wrong content: got %v want %v",
			message.Content, "titkos üzenet")
	}
	if message.UnlocksAt != currenttime {
		t.Errorf("handler returned wrong unlocks_at: got %v want %v",
			message.UnlocksAt, currenttime)
	}
}

func TestCreateMessageThenCheckDb(t *testing.T) {
	currenttime := time.Now().Unix()
	rr := helpCreateMessage("titkos üzenet", currenttime)
	var message models.Message
	json.Unmarshal(rr.Body.Bytes(), &message)

	id, _ := primitive.ObjectIDFromHex(message.ID.Hex())
	filter := bson.M{"_id": id}
	var dbMessage *models.Message
	err := messageCollection.FindOne(context.Background(), filter).Decode(&dbMessage)
	if err != nil || dbMessage.Content != "titkos üzenet" || dbMessage.UnlocksAt != currenttime || dbMessage.User != message.User {
		t.Errorf("mismatch between dbMessage %v and created message %v",
			dbMessage, message)
	}
}
