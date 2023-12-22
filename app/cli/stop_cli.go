package cli

import (
	"fmt"
	"liars-lie/common"
	"net/http"
	"sync"
)

func Stop() {
	endpoints := common.ReadAgentsConfig()
	wg := sync.WaitGroup{}
	for _, e := range endpoints {
		wg.Add(1)
		go stopAgent(fmt.Sprintf("%v/stop", e), &wg)
	}
	wg.Wait()
	common.DeleteAgentsConfig()
}

func stopAgent(endpoint string, wg *sync.WaitGroup) {
	defer wg.Done()
	http.Get(endpoint)
}
