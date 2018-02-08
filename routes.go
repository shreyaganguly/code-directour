package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/new", newHandler).Methods("GET")
	return r
}
