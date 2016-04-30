package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func copyFile(w http.ResponseWriter, r *http.Request, fromPath string, toPath string) (err error) {
	srcFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(srcFile, destFile)

	if err != nil {
		return err
	} else {
		responseHandler(w, r, "copy", toPath, nil)
		return nil
	}
	return nil
}

func copyFolder(w http.ResponseWriter, r *http.Request, fromPath string, toPath string) (err error) {
	fileInfo, err := os.Stat(fromPath)
	if err != nil {
		return err
	}

	if !fileInfo.IsDir() {
		return err
	}

	_, err = os.Open(toPath)
	if !os.IsNotExist(err) {
		return err
	}

	err = os.MkdirAll(toPath, fileInfo.Mode())
	if err != nil {
		return err
	}

	entries, err := ioutil.ReadDir(fromPath)

	for _, entry := range entries {
		srcFilePath := fromPath + "/" + entry.Name()
		destFilePath := toPath + "/" + entry.Name()
		if entry.IsDir() {
			copyFolder(w, r, srcFilePath, destFilePath)
			responseHandler(w, r, "copy", toPath, nil)
			return nil
		} else {
			copyFile(w, r, srcFilePath, destFilePath)
			responseHandler(w, r, "copy", toPath, nil)
			return nil
		}
	}
	return nil
}
