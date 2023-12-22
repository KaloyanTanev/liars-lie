package cli

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestExtend(t *testing.T) {
	f, configName := createConfigFile(t)
	defer func() { f.Close(); os.Remove(configName) }()
	f.Write([]byte("[http://test:12001 http://test:12002]"))

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	go Extend(1, 2, 2, 0.5)

	i := 0
	started := false
	for i < 1000 {
		logs := buf.String()
		if strings.Contains(logs, "Agent running on port :12003") &&
			strings.Contains(logs, "Agent running on port :12004") &&
			strings.Contains(logs, "ready") {
			started = true
			break
		}
		i += 1
		time.Sleep(time.Millisecond * 10)
	}
	if !started {
		t.Fatalf("Service did not start")
	}
}
