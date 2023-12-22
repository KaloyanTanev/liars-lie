package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Agent struct {
	router       *mux.Router
	ShutdownReq  chan bool
	honest       bool
	liarValue    int
	expertValue  map[int]int
	mu           sync.Mutex
	networkValue int
	maxLiarValue int
}

func NewAgent(honest bool, networkValue int, maxLiarValue int, port int) (Agent, *http.Server) {
	r := mux.NewRouter()

	liarValue := -1
	if !honest {
		for liarValue < 0 && liarValue != networkValue {
			liarValue = rand.Intn(maxLiarValue)
		}
	}

	agent := Agent{
		router:       r,
		ShutdownReq:  make(chan bool),
		honest:       honest,
		liarValue:    liarValue,
		networkValue: networkValue,
		expertValue:  make(map[int]int),
		maxLiarValue: maxLiarValue,
	}

	r.HandleFunc("/play", agent.PlayHandler).Methods("GET")
	r.HandleFunc("/stop", agent.ShutdownHandler).Methods("GET")
	r.HandleFunc("/status", agent.StatusHandler).Methods("GET")
	r.HandleFunc("/playexpert", agent.PlayExpertHandler).Methods("POST")
	r.HandleFunc("/lieexpert", agent.LieExpertHandler).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("localhost:%v", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return agent, srv
}
