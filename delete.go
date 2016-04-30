package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"time"
)

func deleteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := ps.ByName("path")

	if isFolder(path) == true {
		deleteFile(w, r, path)
	} else {
		deleteFolder(w, r, path)
	}
}

func deleteFolder(w http.ResponseWriter, r *http.Request, path string) {
	err := os.RemoveAll(path)
	if err != nil {
		errorHandler(w, r, "delete", err, path)
	} else {
		message := ResponseObj{"delete", nil, time.Now(), path, ""}
		response, err := json.Marshal(message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func deleteFile(w http.ResponseWriter, r *http.Request, path string) {
	err := os.Remove(path)
	if err != nil {
		errorHandler(w, r, "delete", err, path)
	} else {
		message := ResponseObj{"delete", nil, time.Now(), path, ""}
		response, err := json.Marshal(message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
