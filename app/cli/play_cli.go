package cli

import (
	"fmt"
	"io"
	"liars-lie/common"
	"log"
	"net/http"
	"strconv"
)

func Play() int {
	endpoints := common.ReadAgentsConfig()
	var agentResponses []int
	var chans []chan int
	for idx, e := range endpoints {
		chans = append(chans, make(chan int))
		go callAgent(fmt.Sprintf("%v/play", e), chans[idx])
	}
	for _, ch := range chans {
		res := <-ch
		if res == -1 {
			continue
		}
		agentResponses = append(agentResponses, res)
	}
	log.Printf("Agents responses: %v\n", agentResponses)
	v := common.MostFrequentInteger(agentResponses)
	log.Printf("Network value: %v\n", v)
	return v
}

func callAgent(endpoint string, c chan<- int) {
	resp, err := http.Get(endpoint)
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
