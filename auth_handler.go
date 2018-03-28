package main

import "net/http"

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	user, err := lookupinUser(getUserName(r))
	if err != nil || user != nil {
		renderer.HTML(w, http.StatusOK, "sign_up", "User Name already exists")
	}
	err = user.Save()
	if err != nil {
		renderer.HTML(w, http.StatusOK, "sign_up", "Internal Error. Please try Again!")
	}
	renderer.HTML(w, http.StatusOK, "index", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	renderer.HTML(w, http.StatusOK, "login", "")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// name := r.FormValue("username")
	// pass := r.FormValue("password")
	redirectTarget := "/"
	// if slackUsersMap[name] != "" && pass == "Alertify123" {
	// 	setSession(name, w)
	// 	redirectTarget = "/index"
	// } else {
	// 	errorText = "Wrong UserName or Password"
	// }

	http.Redirect(w, r, redirectTarget, http.StatusFound)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/", http.StatusFound)
}
