package handlers

import (
	"liars-lie/common"
	"net/http"
)

func (agent *Agent) PlayHandler(w http.ResponseWriter, r *http.Request) {
	if agent.honest {
		common.HTTPJSONResponse(w, agent.networkValue, 200)
	} else {
		common.HTTPJSONResponse(w, agent.liarValue, 200)
	}
}
