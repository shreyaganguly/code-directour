package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler).Methods("GET")
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/sign_up", signUpHandler).Methods("GET")
	r.HandleFunc("/saveuser", saveUserHandler).Methods("POST")
	r.HandleFunc("/index", indexHandler).Methods("GET")
	r.HandleFunc("/new", newHandler).Methods("GET")
	r.HandleFunc("/save", saveHandler).Methods("POST")
	return r
}
