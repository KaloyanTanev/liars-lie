package cli

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func createServer(val int, t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/play" {

			response, err := json.Marshal(val)
			if err != nil {
				t.Fatalf("Error on writing response")
			}
			w.WriteHeader(200)
			w.Write(response)
		}
		if r.URL.Path == "/stop" {
			w.WriteHeader(200)
		}
		if r.URL.Path == "/playexpert" {
			w.WriteHeader(200)
			w.Write([]byte(fmt.Sprintf("%v", val)))
		}
	}))
}

func createConfigFile(t *testing.T) (*os.File, string) {
	configName := "./agents.config"
	f, err := os.Create(configName)
	if err != nil {
		t.Fatalf("failed to create config")
	}
	return f, configName
}

func checkServerIsRunning(url string, t *testing.T) {
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("running server with URL %v did not return 200", url)
	}
}
