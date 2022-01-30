package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func prepareRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.Path("/").HandlerFunc(Handler)
	return router
}