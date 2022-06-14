package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)


func helpCreateUser(email, username, password string) *httptest.ResponseRecorder {
	data := url.Values{}
	data.Set("username", username)
	data.Set("email", email)
	data.Set("password", password)
	req, _ := http.NewRequest("POST", "", bytes.NewBufferString(data.Encode()))	

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)

	handler.ServeHTTP(rr, req)

	return rr
}

func helpLogin(email, password string) *httptest.ResponseRecorder {
	data := url.Values{}
	data.Set("email", email)
	data.Set("password", password)
	req, _ := http.NewRequest("POST", "", bytes.NewBufferString(data.Encode()))	

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)

	handler.ServeHTTP(rr, req)

	return rr
}

// Test user #2, not to be used outside of this function
func TestCreateUser(t *testing.T) {
	data := url.Values{}
	data.Set("username", "testuser2")
	data.Set("email", "testemail2@example.com")
	data.Set("password", "testpass2needstobelonger")
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(data.Encode()))	
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestCreateUserWithSameEmail(t *testing.T) {
	rr_first := helpCreateUser("sameemail@example.com", "differentname1", "testpassword_")
	if rr_first.Code != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr_first.Code, http.StatusCreated)
	}
	rr_second := helpCreateUser("sameemail@example.com", "differentname2", "testpassword_")
	if rr_second.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr_second.Code, http.StatusBadRequest)
	}
	if rr_second.Body.String() != "User with this email already exists\n" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr_second.Body.String(), "User with this email already exists\n")
	}
}

func TestUserCreatedWithSameName(t *testing.T) {
	rr_first := helpCreateUser("diffbutsamename1@example.com", "samename", "testpassword_")
	if rr_first.Code != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr_first.Code, http.StatusCreated)
	}
	rr_second := helpCreateUser("diffbutsamename2@example.com", "samename", "testpassword_")
	if rr_second.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr_second.Code, http.StatusBadRequest)
	}
	if rr_second.Body.String() != "User with this username already exists\n" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr_second.Body.String(), "User with this username already exists\n")
	}
}

func TestCreateUserWithComplexPassword(t *testing.T) {
	rr := helpCreateUser("complexpassword@example.com", "userforcomplex", `KbTjQZbz2i4iBX6VSWsdsyZ7W@6Hz9w7joUR8vb35f*PAUm2imHF97pcK2te4gVd`)
	if rr.Code != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusCreated)
	}
}

func TestCreateUserWithLongPassword(t *testing.T) {
	rr := helpCreateUser("longpassword@example.com", "uforlong", `%K7eJQcX$XX^y2fEvfADry4t*tCiXozoj9xoc3QBa$nTavh3X4T8MKZ^N5Pv3q9G*tZR&EKR*yiqkdL%pWnv25#JLyYDMAm&Ubn!ERaG7YPV%Q5BiWh!L#oc96t8io4y`)
	if rr.Code != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusCreated)
	}
}
