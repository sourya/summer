package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func copyFile(w http.ResponseWriter, r *http.Request, fromPath string, toPath string) int {
	srcFile, err := os.Open(fromPath)
	if err != nil {
		return 1028 // Error reading source file
	}
	defer srcFile.Close()

	destFile, err := os.Create(toPath)
	if err != nil {
		return 1029 // Error creating destination file
	}
	defer destFile.Close()

	_, err = io.Copy(srcFile, destFile)

	if err != nil {
		return 1030 // Error copying file
	} else {
		responseHandler(w, r, "copy", toPath, nil)
		return 0
	}
	return 0
}

func copyFolder(w http.ResponseWriter, r *http.Request, fromPath string, toPath string) int {
	fileInfo, err := os.Stat(fromPath)
	if err != nil {
		return 1031 // Error reading folder stats
	}

	if !fileInfo.IsDir() {
		return 1032 // Is a file
	}

	_, err = os.Open(toPath)
	if !os.IsNotExist(err) {
		return 1033 // Error reading destination folder
	}

	err = os.MkdirAll(toPath, fileInfo.Mode())
	if err != nil {
		return 1034 // Error creating destination folder
	}

	entries, err := ioutil.ReadDir(fromPath)

	for _, entry := range entries {
		srcFilePath := fromPath + "/" + entry.Name()
		destFilePath := toPath + "/" + entry.Name()
		if entry.IsDir() {
			copyFolder(w, r, srcFilePath, destFilePath)
			responseHandler(w, r, "copy", toPath, nil)
			return 0
		} else {
			copyFile(w, r, srcFilePath, destFilePath)
			responseHandler(w, r, "copy", toPath, nil)
			return 0
		}
	}
	return 0
}
