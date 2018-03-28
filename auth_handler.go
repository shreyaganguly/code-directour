package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	renderer.HTML(w, http.StatusOK, "login", "")
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	renderer.HTML(w, http.StatusOK, "sign_up", "")
}

//TODO: common out view
func saveUserHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("username")
	pass := r.FormValue("password")
	user, err := lookupinUser(name)
	if user != nil {
		renderer.HTML(w, http.StatusOK, "sign_up", "User Name already exists")
		return
	}
	if getUserName(r) == "" {
		err = NewUser(name, pass).Save()
		if err != nil {
			renderer.HTML(w, http.StatusOK, "sign_up", "Internal Error. Please try Again!")
			return
		}
		setSession(name, w)
	}
	http.Redirect(w, r, "/index", http.StatusFound)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("username")
	pass := r.FormValue("password")
	fmt.Println(name, pass)
	user, err := lookupinUser(name)
	fmt.Printf("%#v", user)
	if err != nil || user == nil || user.Password != pass {
		renderer.HTML(w, http.StatusOK, "login", "Wrong Username/ Password provided")
		return
	}
	setSession(name, w)
	http.Redirect(w, r, "/index", http.StatusFound)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/", http.StatusFound)
}
