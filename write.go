package main

import (
	"bytes"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func writer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := ps.ByName("path")

	if isFolder(path) == true {
		writeFolder(w, r, path)
	} else {
		buffer := new(bytes.Buffer)
		buffer.ReadFrom(r.Body)
		requestData := buffer.Bytes()

		var requestObj map[string]string
		err := json.Unmarshal(requestData, &requestObj)

		for _, value := range requestObj {
			switch value {
			case "content":
				writeFile(w, r, path, requestObj["content"])
				break
			default:
				errorHandler(w, r, "write", err, path)
				break
			}
		}
	}
}

func writeFolder(w http.ResponseWriter, r *http.Request, path string) {
	err := os.Mkdir(path, 0755)
	if err != nil {
		errorHandler(w, r, "write", err, path)
	}

	message := ResponseObj{"write", nil, time.Now(), path, []string{}}
	response, err := json.Marshal(message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func writeFile(w http.ResponseWriter, r *http.Request, path string, content string) {
	contentByte := []byte(content)
	err := ioutil.WriteFile(path, contentByte, 0755)
	if err != nil {
		errorHandler(w, r, "write", err, path)
	}

	message := ResponseObj{"write", nil, time.Now(), path, []string{}}
	response, err := json.Marshal(message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
