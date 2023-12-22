package handlers

import (
	"testing"
)

func TestShutdownHandler(t *testing.T) {
	a, rec, req := newAgentRecReq(false, t)
	go a.ShutdownHandler(rec, req)

	shutdown := <-a.ShutdownReq

	res := rec.Result()
	if res.StatusCode != 200 {
		t.Fatalf("status code not 200")
	}
	if !shutdown {
		t.Fatalf("did not shutdown")
	}
}
