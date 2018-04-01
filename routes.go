package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler).Methods("GET")
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/logout", logoutHandler).Methods("POST")
	r.HandleFunc("/sign_up", signUpHandler).Methods("GET")
	r.HandleFunc("/saveuser", saveUserHandler).Methods("POST")
	r.HandleFunc("/all", authenticationMiddleware(snippetsHandler)).Methods("GET")
	r.HandleFunc("/shared", authenticationMiddleware(shareListHandler)).Methods("GET")
	r.HandleFunc("/index", authenticationMiddleware(indexHandler)).Methods("GET")
	r.HandleFunc("/new", authenticationMiddleware(newHandler)).Methods("GET")
	r.HandleFunc("/edit/{key}", authenticationMiddleware(editHandler)).Methods("GET")
	r.HandleFunc("/delete/{key}", authenticationMiddleware(deleteHandler)).Methods("GET")
	r.HandleFunc("/share/{key}", authenticationMiddleware(shareHandler)).Methods("POST")
	r.HandleFunc("/link/{key}", authenticationMiddleware(linkHandler)).Methods("GET")
	r.HandleFunc("/save", authenticationMiddleware(saveHandler)).Methods("POST")
	return r
}
