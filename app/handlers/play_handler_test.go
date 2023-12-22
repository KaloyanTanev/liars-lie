package handlers

import (
	"io"
	"strconv"
	"testing"
)

func TestPlayHandlerLiar(t *testing.T) {
	a, rec, req := newAgentRecReq(false, t)
	a.PlayHandler(rec, req)

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
	if resInt != a.liarValue {
		t.Fatalf("res not equal to liar value")
	}
}

func TestPlayHandlerHonest(t *testing.T) {
	a, rec, req := newAgentRecReq(true, t)
	a.PlayHandler(rec, req)

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
		t.Fatalf("res not equal to liar value")
	}
}
