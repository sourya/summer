package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

func deleteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := ps.ByName("path")

	if isFolder(path) == true {
		err := deleteFile(w, r, path)
		if err != nil {
			errorHandler(w, r, "delete", err, path)
		}
	} else {
		err := deleteFolder(w, r, path)
		if err != nil {
			errorHandler(w, r, "delete", err, path)
		}
	}
}

func deleteFolder(w http.ResponseWriter, r *http.Request, path string) (err error) {
	err = os.RemoveAll(path)
	if err != nil {
		return err
	} else {
		responseHandler(w, r, "delete", path, nil)
	}
	return nil
}

func deleteFile(w http.ResponseWriter, r *http.Request, path string) (err error) {
	err = os.Remove(path)
	if err != nil {
		return err
	} else {
		responseHandler(w, r, "delete", path, nil)
	}
	return nil
}
