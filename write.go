package main

import (
	"bytes"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"os"
)

func writeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := ps.ByName("path")

	if isFolder(path) == true {
		err := writeFolder(w, r, path)
		if err != nil {
			errorHandler(w, r, "write", err, path)
		}
	} else {
		buffer := new(bytes.Buffer)
		buffer.ReadFrom(r.Body)
		requestData := buffer.Bytes()

		if len(requestData) == 0 {
			err := writeFile(w, r, path, "")
			if err != nil {
				errorHandler(w, r, "write", err, path)
			}
		} else {
			bodyParser(requestData, r, w, path)
		}
	}
}

func writeFolder(w http.ResponseWriter, r *http.Request, path string) (err error) {
	err = os.Mkdir(path, 0755)
	if err != nil {
		return err
	} else {
		responseHandler(w, r, "write", path, nil)
		return nil
	}
	return nil
}

func writeFile(w http.ResponseWriter, r *http.Request, path string, content string) (err error) {
	contentByte := []byte(content)
	err = ioutil.WriteFile(path, contentByte, 0755)
	if err != nil {
		return err
	} else {
		responseHandler(w, r, "write", path, nil)
		return nil
	}
	return nil
}
