package handlers

import (
	"liars-lie/common"
	"net/http"
)

func (agent *Agent) StatusHandler(w http.ResponseWriter, r *http.Request) {
	common.HTTPEmptyResponse(w, 200)
}
