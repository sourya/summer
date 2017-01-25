package main

import (
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func readHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := setRoot(ps.ByName("path"))

	if isFolder(path) == true {
		errCode := readFolder(w, r, path)
		if errCode != 0 {
			errorHandler(w, r, "read", errCode, path)
		}
	} else {
		errCode := readFile(w, r, path)
		if errCode != 0 {
			errorHandler(w, r, "read", errCode, path)
		}
	}
}

func readFolder(w http.ResponseWriter, r *http.Request, path string) int {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return 1024 // Error reading directory
	} else {
		dirList := []string{}

		for _, file := range files {
			dirList = append(dirList, file.Name())
		}

		responseHandler(w, r, "read", path, dirList)
		return 0
	}
}

func readFile(w http.ResponseWriter, r *http.Request, path string) int {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return 1025 // Error reading file
	} else {
		responseHandler(w, r, "read", path, string(content))
		return 0
	}
}
