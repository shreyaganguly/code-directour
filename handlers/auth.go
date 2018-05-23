package handlers

import (
	"net/http"
	"strings"

	"github.com/shreyaganguly/code-directour/db"
	"github.com/shreyaganguly/code-directour/models"
	"github.com/shreyaganguly/code-directour/util"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	util.Renderer.HTML(w, http.StatusOK, "login", "")
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	util.Renderer.HTML(w, http.StatusOK, "sign_up", "")
}

func saveUserHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("username")
	pass := r.FormValue("password")
	user, _ := db.LookupinUser(strings.ToLower(name))
	if user != nil {
		util.Renderer.HTML(w, http.StatusOK, "sign_up", "User Name already exists")
		return
	}
	err := db.Update(models.NewUser(strings.ToLower(name), pass))
	if err != nil {
		util.Renderer.HTML(w, http.StatusOK, "sign_up", "Internal Error. Please try Again!")
		return
	}
	util.SetSession(name, w)
	http.Redirect(w, r, "/index", http.StatusFound)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("username")
	pass := r.FormValue("password")
	user, err := db.LookupinUser(strings.ToLower(name))
	if err != nil || user == nil || user.Password != pass {
		util.Renderer.HTML(w, http.StatusOK, "login", "Wrong Username/ Password provided")
		return
	}
	util.SetSession(name, w)
	http.Redirect(w, r, "/index", http.StatusFound)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	util.ClearSession(w)
	http.Redirect(w, r, "/", http.StatusFound)
}
