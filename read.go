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
	} else {
		dirList := []string{}

		for _, file := range files {
			dirList = append(dirList, file.Name())
		}

		message := ResponseObj{"read", nil, time.Now(), path, dirList}
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

func readFile(w http.ResponseWriter, r *http.Request, path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		errorHandler(w, r, "read", err, path)
	} else {
		message := ResponseObj{"read", nil, time.Now(), path, string(content)}
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
