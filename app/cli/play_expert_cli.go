package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"liars-lie/app/handlers"
	"liars-lie/common"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func PlayExpert(numberOfAgents int, liarRatio float64) int {
	endpoints := common.ReadAgentsConfig()
	// get random agents equal to numberOfAgents
	var endpointsToQuery []string
	for i := 0; i < numberOfAgents; i++ {
		randBelow := len(endpoints)
		j := rand.Intn(randBelow)
		endpointsToQuery = append(endpointsToQuery, endpoints[j])
		endpoints = common.RemoveIndex(endpoints, j)
	}

	var agentResponses []int
	var chans []chan int
	for idx, e := range endpointsToQuery {
		chans = append(chans, make(chan int))
		req := handlers.PlayExpertRequest{
			Id:        idx,
			Endpoints: endpointsToQuery,
		}
		go callAgentExpert(fmt.Sprintf("%v/playexpert", e), req, chans[idx])
	}
	for _, ch := range chans {
		res := <-ch
		if res == -1 {
			continue
		}
		agentResponses = append(agentResponses, res)
	}
	log.Printf("Agents responses: %v\n", agentResponses)
	// get all responses, their ratio compared to the total and get the closest one to liarRatio
	distinct := common.GetDistinctIntegers(agentResponses)
	lenResponses := float64(len(agentResponses))
	closest := 1.0
	closestResponse := -1
	honestRatio := 1.0 - liarRatio
	for k, v := range distinct {
		distanceToPredictedRatio := math.Abs(honestRatio - (float64(v) / lenResponses))
		if distanceToPredictedRatio < closest {
			closest = distanceToPredictedRatio
			closestResponse = k
		}
	}
	log.Printf("Network value: %v\n", closestResponse)
	return closestResponse
}

func callAgentExpert(endpoint string, req handlers.PlayExpertRequest, c chan<- int) {
	reqJSON, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		log.Printf("Agent %v is killed", endpoint)
		c <- -1
		return
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("%v", err.Error())
		c <- -1
		return
	}
	res, err := strconv.Atoi(string(respBody))
	if err != nil {
		log.Fatalf("%v", err.Error())
		c <- -1
		return
	}
	c <- res
}
