package main

import (
	"net/http"
	"os"
)

func linkHandler(w http.ResponseWriter, r *http.Request, fromPath string, toPath string) int {
	err := os.Symlink(fromPath, toPath)
	if err != nil {
		return 1038 // Error creating symlink
	} else {
		responseHandler(w, r, "symlink", toPath, 0)
	}
	return 0
}
