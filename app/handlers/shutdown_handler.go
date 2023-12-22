package handlers

import (
	"liars-lie/common"
	"net/http"
)

func (agent *Agent) ShutdownHandler(w http.ResponseWriter, r *http.Request) {
	common.HTTPEmptyResponse(w, 200)
	agent.ShutdownReq <- true
}
