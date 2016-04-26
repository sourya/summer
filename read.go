package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"time"
)

func reader(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := ps.ByName("path")

	if isFolder(path) == true {
		readFolder(w, r, path)
	} else {
		readFile(w, r, path)
	}
}

func readFolder(w http.ResponseWriter, r *http.Request, path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		errorHandler(w, r, "read", err, path)
	}

	dirList := []string{}

	for _, file := range files {
		dirList = append(dirList, file.Name())
	}

	message := ResponseObj{"read", nil, time.Now(), path, dirList}
	response, err := json.Marshal(message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func readFile(w http.ResponseWriter, r *http.Request, path string) {
	_, err := ioutil.ReadFile(path)
	if err != nil {
		errorHandler(w, r, "read", err, path)
	}

	message := ResponseObj{"read", nil, time.Now(), path, []string{}}
	response, err := json.Marshal(message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
