package cli

import (
	"fmt"
	"os"
	"testing"
)

func TestPlayExpert(t *testing.T) {
	testServer1 := createServer(5, t)
	defer testServer1.Close()
	testServer2 := createServer(5, t)
	defer testServer2.Close()
	testServer3 := createServer(10, t)
	defer testServer3.Close()

	f, configName := createConfigFile(t)
	defer func() { f.Close(); os.Remove(configName) }()

	urls := []string{testServer1.URL, testServer2.URL, testServer3.URL}
	_, err := f.Write([]byte(fmt.Sprintf("%v", urls)))
	if err != nil {
		t.Fatalf("failed to write to config")
	}

	v := PlayExpert(3, 0.33)
	if v != 5 {
		t.Fatalf("network value is not 5")
	}
}
