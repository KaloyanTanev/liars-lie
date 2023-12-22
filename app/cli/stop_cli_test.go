package cli

import (
	"fmt"
	"net/http/httptest"
	"os"
	"testing"
)

func TestStop(t *testing.T) {
	numServers := 2
	servers := []*httptest.Server{}
	for i := 0; i < numServers; i++ {
		testServer := createServer(0, t)
		defer testServer.Close()
		servers = append(servers, testServer)
	}
	f, configName := createConfigFile(t)
	defer func() { f.Close(); os.Remove(configName) }()

	urls := []string{}
	for _, s := range servers {
		urls = append(urls, s.URL)
	}
	_, err := f.Write([]byte(fmt.Sprintf("%v", urls)))
	if err != nil {
		t.Fatalf("%v", err)
	}

	for _, url := range urls {
		checkServerIsRunning(url, t)
	}

	Stop()
}
