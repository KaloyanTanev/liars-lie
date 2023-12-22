package cli

import (
	"fmt"
	"liars-lie/common"
	"net/http"
)

func Kill(id int) {
	endpoints := common.ReadAgentsConfig()
	killEndpoint := endpoints[id]
	http.Get(fmt.Sprintf("%v/stop", killEndpoint))
}
