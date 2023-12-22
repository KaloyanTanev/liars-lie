package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLieExpertHandlerLiar(t *testing.T) {
	a, _ := NewAgent(false, 1, 5, 8000)
	rec := httptest.NewRecorder()
	reqBody := LieExpertRequest{Id: 1, Value: 5}
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("%v", err.Error())
	}
	reader := bytes.NewReader(reqBodyBytes)
	req, err := http.NewRequest("GET", "http://test.com", reader)
	if err != nil {
		t.Fatalf("%v", err.Error())
	}
	a.LieExpertHandler(rec, req)

	if rec.Result().StatusCode != 200 {
		t.Fatalf("status code not 200")
	}

	if a.expertValue[1] != 5 {
		t.Fatalf("proposed value not updated")
	}
}
func TestLieExpertHandlerHonest(t *testing.T) {
	a, _ := NewAgent(true, 1, 0, 8000)
	rec := httptest.NewRecorder()
	reqBody := LieExpertRequest{Id: 1, Value: 5}
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("%v", err.Error())
	}
	reader := bytes.NewReader(reqBodyBytes)
	req, err := http.NewRequest("GET", "http://test.com", reader)
	if err != nil {
		t.Fatalf("%v", err.Error())
	}
	a.LieExpertHandler(rec, req)

	if rec.Result().StatusCode != 400 {
		t.Fatalf("status code not 400")
	}
}
