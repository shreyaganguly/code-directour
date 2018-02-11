package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/new", newHandler).Methods("GET")
	r.HandleFunc("/list", listHandler).Methods("GET")
	r.HandleFunc("/save", saveHandler).Methods("POST")
	return r
}
