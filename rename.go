package main

import (
	"net/http"
	"os"
)

func renameHandler(w http.ResponseWriter, r *http.Request, fromPath string, toPath string) int {
	err := os.Rename(fromPath, toPath)
	if err != nil {
		return 1036 // Error renaming file/folder
	} else {
		responseHandler(w, r, "rename", toPath, 0)
	}
	return 0
}
