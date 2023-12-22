package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"liars-lie/common"
	"log"
	"net/http"
	"time"
)

type PlayExpertRequest struct {
	Id        int
	Endpoints []string
}

func (agent *Agent) PlayExpertHandler(w http.ResponseWriter, r *http.Request) {
	// If agent is honest return network value and skip any other steps.
	if agent.honest {
		common.HTTPJSONResponse(w, agent.networkValue, 200)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req PlayExpertRequest
	err := decoder.Decode(&req)
	if err != nil {
		log.Fatal(err)
	}

	// Call other agents to propose your lying value.
	nOfLiarsProposedTo := agent.proposeToOtherLiars(req.Id, agent.liarValue, req.Endpoints)
	i := 0
	// Wait for as many liars to propose their value, as much as the current liar proposed to with timeout 5 sec.
	for len(agent.expertValue) < nOfLiarsProposedTo && i < 50 {
		time.Sleep(time.Millisecond * 100)
		i += 1
	}

	// Choose the value from the smallest index (input from play_exper_cli.go for loop).
	// As the slice of agents is always completely randomized, this will always be a random agent.
	// This ensures all agents agree on the same value (except in the case of timeout from the for loop above).
	keys := common.MapKeys(agent.expertValue)
	firstProposer := common.SmallestInteger(keys)
	expertLiarValue := agent.expertValue[firstProposer]
	agent.expertValue = make(map[int]int)
	common.HTTPJSONResponse(w, expertLiarValue, 200)
}

func (agent *Agent) proposeToOtherLiars(id int, proposedValue int, endpoints []string) int {
	res := LieExpertRequest{
		Id:    id,
		Value: proposedValue,
	}
	resJSON, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	// Track how many agents accepted our proposal.
	numMocked := 0
	for _, e := range endpoints {
		r, err := http.Post(fmt.Sprintf("%v/lieexpert", e), "application/json", bytes.NewBuffer(resJSON))
		if err != nil {
			log.Printf("Agent %v is killed", e)
			continue
		}
		if r.StatusCode > 400 {
			log.Fatal("Inter-agent communicatoin returned HTTP error code")
		}
		if r.StatusCode == 200 {
			numMocked += 1
		}
	}
	return numMocked
}
