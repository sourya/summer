package main

import (
	"net/http"
	"os"
)

func renameHandler(w http.ResponseWriter, r *http.Request, fromPath string, toPath string) (err error) {
	err = os.Rename(fromPath, toPath)
	if err != nil {
		return err
	} else {
		responseHandler(w, r, "rename", toPath, nil)
	}
	return nil
}
