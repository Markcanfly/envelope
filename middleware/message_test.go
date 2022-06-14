package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

// TODO add message creation and viewing tests
