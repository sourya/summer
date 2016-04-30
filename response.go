package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func responseDispatcher(w http.ResponseWriter, r *http.Request, message ResponseObj, isErr bool) {
	response, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if isErr == true {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write(response)
}

func responseHandler(w http.ResponseWriter, r *http.Request, operation string, path string, content interface{}) {
	message := ResponseObj{operation, nil, time.Now(), path, content}
	responseDispatcher(w, r, message, false)
}

func errorHandler(w http.ResponseWriter, r *http.Request, operation string, err error, path string) {
	message := ResponseObj{operation, err, time.Now(), path, ""}
	responseDispatcher(w, r, message, true)
}
