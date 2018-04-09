package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shreyaganguly/code-directour/interceptors"
)

func SetUpRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler).Methods("GET")
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/logout", logoutHandler).Methods("POST")
	r.HandleFunc("/sign_up", signUpHandler).Methods("GET")
	r.HandleFunc("/saveuser", saveUserHandler).Methods("POST")
	r.HandleFunc("/all", interceptors.AuthenticationMiddleware(snippetsHandler)).Methods("GET")
	r.HandleFunc("/shared", interceptors.AuthenticationMiddleware(shareListHandler)).Methods("GET")
	r.HandleFunc("/index", interceptors.AuthenticationMiddleware(indexHandler)).Methods("GET")
	r.HandleFunc("/new", interceptors.AuthenticationMiddleware(newHandler)).Methods("GET")
	r.HandleFunc("/edit/{key}", interceptors.AuthenticationMiddleware(editHandler)).Methods("GET")
	r.HandleFunc("/delete/{key}", interceptors.AuthenticationMiddleware(deleteHandler)).Methods("GET")
	r.HandleFunc("/share/{key}", interceptors.AuthenticationMiddleware(shareHandler)).Methods("POST")
	r.HandleFunc("/sharemail/{key}", interceptors.AuthenticationMiddleware(shareEmailHandler)).Methods("POST")
	r.HandleFunc("/link/{key}/{name}", linkHandler).Methods("GET")
	r.HandleFunc("/save", interceptors.AuthenticationMiddleware(saveHandler)).Methods("POST")
	return r
}
