package main

import (
	"bytes"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"os"
)

func writeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := setRoot(ps.ByName("path"))

	if isFolder(path) == true {
		errCode := writeFolder(w, r, path)
		if errCode != 0 {
			errorHandler(w, r, "write", errCode, path)
		}
	} else {
		buffer := new(bytes.Buffer)
		buffer.ReadFrom(r.Body)
		requestData := buffer.Bytes()

		if len(requestData) == 0 {
			errCode := writeFile(w, r, path, "")
			if errCode != 0 {
				errorHandler(w, r, "write", errCode, path)
			}
		} else {
			bodyParser(requestData, r, w, path)
		}
	}
}

func writeFolder(w http.ResponseWriter, r *http.Request, path string) int {
	err := os.Mkdir(path, 0755)
	if err != nil {
		return 1026 // Error creating folder
	} else {
		responseHandler(w, r, "write", path, 0)
		return 0
	}
	return 0
}

func writeFile(w http.ResponseWriter, r *http.Request, path string, content string) int {
	contentByte := []byte(content)
	err := ioutil.WriteFile(path, contentByte, 0755)
	if err != nil {
		return 1027 // Error creating/writing file
	} else {
		responseHandler(w, r, "write", path, 0)
		return 0
	}
	return 0
}
