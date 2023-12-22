package cli

import (
	"context"
	"fmt"
	"liars-lie/app/handlers"
	"liars-lie/common"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func Start(networkValue int, maxLiarValue int, numberOfAgents int, liarRatio float64) {
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

	startPort := 13001

	var endpoints []string
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
	defer common.DeleteAgentsConfig()

	log.Println("ready")
	// if all agents in this process are killed, end the process
	wg.Wait()
}

func runAgent(honest bool, networkValue int, maxLiarValue int, port int, wg *sync.WaitGroup, c chan<- bool) {
	defer wg.Done()

	agent, srv := handlers.NewAgent(honest, networkValue, maxLiarValue, port)

	go func() {
		log.Printf("Agent running on port :%v\n", port)
		c <- true
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server boot up error %v", err)
		}
	}()

	// wait for /stop request
	<-agent.ShutdownReq
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Agent on port %v shut down gracefully\n", port)
}

func testAgent(port int, c chan<- bool) {
	i := 0
	// test agent on /status endpoint each second, timeout 10 seconds
	for i < 10 {
		resp, err := http.Get(fmt.Sprintf("http://localhost:%v/status", port))
		if err == nil && resp.StatusCode == 200 {
			c <- true
			return
		}
		i += 1
		time.Sleep(time.Second)
	}
	c <- false
}
