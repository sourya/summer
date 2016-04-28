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

		if len(requestData) == 0 {
			writeFile(w, r, path, "")
		} else {
			var requestObj map[string]string
			err := json.Unmarshal(requestData, &requestObj)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			for key, value := range requestObj {
				switch key {
				case "content":
					writeFile(w, r, path, value)
					break
				default:
					errorHandler(w, r, "write", err, path)
					break
				}
			}
		}
	}
}

func writeFolder(w http.ResponseWriter, r *http.Request, path string) {
	err := os.Mkdir(path, 0755)
	if err != nil {
		errorHandler(w, r, "write", err, path)
	} else {
		message := ResponseObj{"write", nil, time.Now(), path, ""}
		response, err := json.Marshal(message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func writeFile(w http.ResponseWriter, r *http.Request, path string, content string) {
	contentByte := []byte(content)
	err := ioutil.WriteFile(path, contentByte, 0755)
	if err != nil {
		errorHandler(w, r, "write", err, path)
	} else {
		message := ResponseObj{"write", nil, time.Now(), path, ""}
		response, err := json.Marshal(message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

}
