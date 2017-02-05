package main

import (
	"github.com/gorilla/handlers"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

func authenticator(router http.Handler) http.Handler {
    access_key_authenticator := func(w http.ResponseWriter, r *http.Request) {
        access_key := viper.Get("access_key").(string)
        entered_key := r.URL.Query().Get("access_key")

        if entered_key == access_key {
           router.ServeHTTP(w, r)
        } else {
            w.WriteHeader(401)
        }
    }

    return http.HandlerFunc(access_key_authenticator)
}

func middleware(router http.Handler) http.Handler {
	// Logger
	if (viper.GetBool("logger")) == true {
		return handlers.CombinedLoggingHandler(os.Stdout, router)
	}

	return router
}
