package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/lieexpert" {
			w.WriteHeader(200)
		}
	}))
}

func newAgentRecReq(honest bool, t *testing.T) (Agent, *httptest.ResponseRecorder, *http.Request) {
	a, _ := NewAgent(honest, 1, 5, 8000)

	rec := httptest.NewRecorder()
	reader := bytes.NewReader([]byte{})
	req, err := http.NewRequest("GET", "http://test.com", reader)
	if err != nil {
		t.Fatalf("%v", err.Error())
	}

	return a, rec, req
}
