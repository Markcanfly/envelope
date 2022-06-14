package middleware

import (
	"bytes"
	"encoding/json"
	"envelope/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func helpCreateMessage(content string, timestamp int64) *httptest.ResponseRecorder {
	message := fmt.Sprintf(`{"content":"%s", "unlocks_at": %d}`, content, timestamp)
	req, _ := http.NewRequest("POST", "", bytes.NewBufferString(message))	
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateMessage)

	handler.ServeHTTP(rr, req)
	return rr
}

func TestGetOpenedMessagesEmpty(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)	
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllOpenedMessages)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	if rr.Body.String() != "null\n" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "[]")
	}
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
			status, http.StatusOK)
	}
}

func TestCreateMessage(t *testing.T) {
	message := fmt.Sprintf(`{"content":"titkos üzenet", "unlocks_at": %d}`, time.Now().Unix())
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(message))	
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
