package common

import (
	"encoding/json"
	"log"
	"net/http"
)

func HTTPJSONResponse(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(data)
	if err != nil {
		log.Panicln(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(response)
}

func HTTPEmptyResponse(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}
