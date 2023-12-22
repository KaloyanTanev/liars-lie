package cli

import (
	"fmt"
	"liars-lie/common"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"sync"
)

func Extend(networkValue int, maxLiarValue int, numberOfAgents int, liarRatio float64) {
	// randomize the order at which honest and liar agents are created
	honest := int(float64(numberOfAgents) * (1 - liarRatio))
	liars := numberOfAgents - honest
	var agentsArray []bool
	for i := 0; i < liars; i++ {
		agentsArray = append(agentsArray, false)
	}
	for i := 0; i < honest; i++ {
		agentsArray = append(agentsArray, true)
	}
	for i := range agentsArray {
		j := rand.Intn(i + 1)
		agentsArray[i], agentsArray[j] = agentsArray[j], agentsArray[i]
	}

	endpoints := common.ReadAgentsConfig()
	// get the last port used
	startPort, err := strconv.Atoi(strings.Split(endpoints[len(endpoints)-1], ":")[2])
	if err != nil {
		log.Fatal(err)
	}
	startPort += 1
	wg := sync.WaitGroup{}
	var runChans []chan bool
	for idx, agent := range agentsArray {
		runChans = append(runChans, make(chan bool))
		wg.Add(1)
		port := startPort + idx
		go runAgent(agent, networkValue, maxLiarValue, port, &wg, runChans[idx])
		endpoints = append(endpoints, fmt.Sprintf("http://localhost:%v", port))
	}

	// wait for all agents to be created
	var statusChans []chan bool
	for idx, ch := range runChans {
		<-ch
		port := startPort + idx
		statusChans = append(statusChans, make(chan bool))
		go testAgent(port, statusChans[idx])
	}

	// check status of all servers
	for idx, ch := range statusChans {
		res := <-ch
		port := startPort + idx
		if !res {
			log.Fatalf("Agent on port %v could not start", port)
		}
	}

	common.OverwriteAgentsConfig(endpoints)

	log.Println("ready")
	// if all agents in this process are killed, end the process
	wg.Wait()
}
