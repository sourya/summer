package main

import (
	"bytes"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func modifyHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	toPath := ps.ByName("path")

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(r.Body)
	requestData := buffer.Bytes()

	if len(requestData) == 0 {
		errorHandler(w, r, "modify", 1035, toPath) // Request body not found
	} else {
		parsedKey, parsedValue := bodyParser(requestData, r, w, toPath)
		switch parsedKey {
		case "renameFrom":
			errCode := renameHandler(w, r, parsedValue, toPath)
			if errCode != 0 {
				errorHandler(w, r, "copy", errCode, toPath)
			}
			break
		case "copyFrom":
			if isFolder(toPath) == true {
				errCode := copyFolder(w, r, parsedValue, toPath)
				if errCode != 0 {
					errorHandler(w, r, "copy", errCode, toPath)
				}
			} else {
				errCode := copyFile(w, r, parsedValue, toPath)
				if errCode != 0 {
					errorHandler(w, r, "copy", errCode, toPath)
				}
			}
		}
	}
}
