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
		errorHandler(w, r, "modify", nil, toPath)
	} else {
		parsedKey, parsedValue := bodyParser(requestData, r, w, toPath)
		switch parsedKey {
		case "renameFrom":
			err := renameHandler(w, r, parsedValue, toPath)
			if err != nil {
				errorHandler(w, r, "copy", err, toPath)
			}
			break
		case "copyFrom":
			if isFolder(toPath) == true {
				err := copyFolder(w, r, parsedValue, toPath)
				if err != nil {
					errorHandler(w, r, "copy", err, toPath)
				}
			} else {
				err := copyFile(w, r, parsedValue, toPath)
				if err != nil {
					errorHandler(w, r, "copy", err, toPath)
				}
			}
		}
	}
}
