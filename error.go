package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func errorHandler(w http.ResponseWriter, r *http.Request, operation string, err error, path string) {
	message := ResponseObj{operation, err, time.Now(), path, []string{}}
	response, err := json.Marshal(message)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(response)
}
