package main

import (
	"encoding/json"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

func setRoot(path string) string {
	return viper.Get("root").(string) + path
}

func isFolder(path string) bool {
	return strings.HasSuffix(path, "/")
}

func bodyParser(requestData []byte, r *http.Request, w http.ResponseWriter, path string) (key string, value string) {
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
			return "", ""
		case "renameFrom":
			if len(value) != 0 {
				return key, value
			}
			break
		case "copyFrom":
			if len(value) != 0 {
				return key, value
			}
			break
		default:
			errorHandler(w, r, "write", 1037, path) // Unsupported keys in request body
			break
		}
	}
	return
}
