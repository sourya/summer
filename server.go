package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	_ = viper.ReadInConfig()
}

func main() {
	router := httprouter.New()

	router.GET("/:path", Read)
	router.PUT("/:path", Write)
	router.DELETE("/:path", Detele)
	router.POST("/:path", Modify)

	log.Fatal(http.ListenAndServe(":"+viper.Get("appPort").(string), router))
}

func isFolder(path string) bool {
	return strings.HasSuffix(path, "/")
}
