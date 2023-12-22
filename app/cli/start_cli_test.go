package cli

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	go Start(1, 2, 2, 0.5)

	i := 0
	started := false
	for i < 1000 {
		logs := buf.String()
		if strings.Contains(logs, "Agent running on port :12002") &&
			strings.Contains(logs, "Agent running on port :12001") &&
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
