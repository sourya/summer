package main

import (
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func readHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := ps.ByName("path")

	if isFolder(path) == true {
		err := readFolder(w, r, path)
		if err != nil {
			errorHandler(w, r, "read", err, path)
		}
	} else {
		err := readFile(w, r, path)
		if err != nil {
			errorHandler(w, r, "read", err, path)
		}
	}
}

func readFolder(w http.ResponseWriter, r *http.Request, path string) (err error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	} else {
		dirList := []string{}

		for _, file := range files {
			dirList = append(dirList, file.Name())
		}

		responseHandler(w, r, "read", path, dirList)
		return nil
	}
}

func readFile(w http.ResponseWriter, r *http.Request, path string) (err error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	} else {
		responseHandler(w, r, "read", path, string(content))
		return nil
	}
	return nil
}
