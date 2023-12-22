package handlers

import (
	"encoding/json"
	"liars-lie/common"
	"log"
	"net/http"
)

type LieExpertRequest struct {
	Id    int
	Value int
}

func (agent *Agent) LieExpertHandler(w http.ResponseWriter, r *http.Request) {
	// If agent is honest, return bad request.
	if agent.honest {
		common.HTTPEmptyResponse(w, 400)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var t LieExpertRequest
	err := decoder.Decode(&t)
	if err != nil {
		log.Fatal(err)
	}

	// Add value of proposer with unique id to the proposed values map.
	// If two or more 'playexpert' actions are executed one after the other,
	// there might be race condition, that's why we use locks.
	agent.mu.Lock()
	agent.expertValue[t.Id] = t.Value
	agent.mu.Unlock()

	common.HTTPEmptyResponse(w, 200)
}
