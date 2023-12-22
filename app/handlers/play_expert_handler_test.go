package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestPlayExpertHandlerLiar(t *testing.T) {
	a, _ := NewAgent(false, 1, 5, 8000)
	a.expertValue[2] = 5

	agent2 := createServer(t)
	defer agent2.Close()

	rec := httptest.NewRecorder()
	rec.Body = bytes.NewBuffer([]byte{})
	reqBody := PlayExpertRequest{Id: 1, Endpoints: []string{agent2.URL}}
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("%v", err.Error())
	}
	reader := bytes.NewReader(reqBodyBytes)
	req, err := http.NewRequest("GET", "http://test.com", reader)
	if err != nil {
		t.Fatalf("%v", err.Error())
	}
	a.PlayExpertHandler(rec, req)

	res := rec.Result()
	if res.StatusCode != 200 {
		t.Fatalf("status code not 200")
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("%v", err.Error())
	}
	resInt, err := strconv.Atoi(string(resBody))
	if err != nil {
		t.Fatalf("%v", err.Error())
	}
	if resInt != 5 {
		t.Fatalf("res not equal to 5")
	}
}

func TestPlayExpertHandlerTestHonest(t *testing.T) {
	a, _ := NewAgent(true, 1, 5, 8000)

	rec := httptest.NewRecorder()
	rec.Body = bytes.NewBuffer([]byte{})
	reqBody := PlayExpertRequest{Id: 1, Endpoints: []string{}}
	reqBodyB, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("%v", err.Error())
	}
	reader := bytes.NewReader(reqBodyB)
	req, err := http.NewRequest("GET", "http://test.com", reader)
	if err != nil {
		t.Fatalf("%v", err.Error())
	}
	a.PlayExpertHandler(rec, req)

	res := rec.Result()
	if res.StatusCode != 200 {
		t.Fatalf("status code not 200")
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("%v", err.Error())
	}
	resInt, err := strconv.Atoi(string(resBody))
	if err != nil {
		t.Fatalf("%v", err.Error())
	}
	if resInt != a.networkValue {
		t.Fatalf("res not equal to network value")
	}
}
