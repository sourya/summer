package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strings"
	"time"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	_ = viper.ReadInConfig()
}

type ResponseObj struct {
	Operation string      `json:"operation"`
	Err       error       `json:"error"`
	Timestamp time.Time   `json:"timestamp"`
	Path      string      `json:"path"`
	Content   interface{} `json:"content"`
}

func main() {
	router := httprouter.New()

	router.GET("/*path", reader)
	router.PUT("/*path", writer)
	// router.DELETE("/:path", detele)
	// router.POST("/:path", modify)

	log.Println("Summer server listening at port" + ":" + viper.Get("appPort").(string))
	log.Fatal(http.ListenAndServe(":"+viper.Get("appPort").(string), router))
}

func isFolder(path string) bool {
	return strings.HasSuffix(path, "/")
}
