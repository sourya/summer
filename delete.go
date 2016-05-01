package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

func deleteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := ps.ByName("path")

	if isFolder(path) == true {
		errCode := deleteFile(w, r, path)
		if errCode != 0 {
			errorHandler(w, r, "delete", errCode, path)
		}
	} else {
		errCode := deleteFolder(w, r, path)
		if errCode != 0 {
			errorHandler(w, r, "delete", errCode, path)
		}
	}
}

func deleteFolder(w http.ResponseWriter, r *http.Request, path string) int {
	err := os.RemoveAll(path)
	if err != nil {
		return 1028 // Error deleting folder
	} else {
		responseHandler(w, r, "delete", path, 0)
	}
	return 0
}

func deleteFile(w http.ResponseWriter, r *http.Request, path string) int {
	err := os.Remove(path)
	if err != nil {
		return 1029 // Error deleting file
	} else {
		responseHandler(w, r, "delete", path, 0)
	}
	return 0
}
