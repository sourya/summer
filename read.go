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
		readFolder(w, r, ps, path)
	} else {
		readFile(w, r, ps, path)
	}
}

func readFolder(w http.ResponseWriter, r *http.Request, ps httprouter.Params, path string) {
	type Response struct {
		Status    string    `json:"status"`
		Err       error     `json:"error"`
		Timestamp time.Time `json:"timestamp"`
		Path      string    `json:"path"`
		Content   []string  `json:"content"`
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		message := Response{"err", err, time.Now(), path, []string{}}
		response, err := json.Marshal(message)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

	dirList := []string{}

	for _, file := range files {
		dirList = append(dirList, file)
	}

	message := Response{"ok", nil, time.Now(), path, dirList}
	response, err := json.Marshal(message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func readFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, path string) {
	type Response struct {
		Status    string    `json:"status"`
		Err       error     `json:"error"`
		Timestamp time.Time `json:"timestamp"`
		Path      string    `json:"path"`
		Content   string    `json:"content"`
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		message := Response{"err", err, time.Now(), path, ""}
		response, err := json.Marshal(message)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

	message := Response{"ok", nil, time.Now(), path, content}
	response, err := json.Marshal(message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
