package main

import (
	"github.com/gorilla/handlers"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

func middleware(router http.Handler) http.Handler {
	// Logger
	if (viper.GetBool("logger")) == true {
		return handlers.CombinedLoggingHandler(os.Stdout, router)
	}
	return router
}
