package handlers

import (
	"testing"
)

func TestStatus(t *testing.T) {
	a, rec, req := newAgentRecReq(false, t)
	a.StatusHandler(rec, req)

	res := rec.Result()
	if res.StatusCode != 200 {
		t.Fatalf("status code not 200")
	}
}
