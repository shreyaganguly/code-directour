package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shreyaganguly/code-directour/db"
	"github.com/shreyaganguly/code-directour/util"
)

func profileHandler(w http.ResponseWriter, r *http.Request) {
	user, err := db.LookupinUser(util.GetUserName(r))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	util.Renderer.HTML(w, http.StatusOK, "profile", user)
}

func profileSaveHandler(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	user, err := db.LookupinUser(util.GetUserName(r))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	switch t := args["type"]; t {
	case "link":
		user.Endpoint = r.PostFormValue("endpoint")
	case "email":
	case "slack":
	default:
		http.Error(w, "Some error occured!!", http.StatusInternalServerError)
	}
	err = db.Update(user)
	if err != nil {
		//TODO: cleaner error handling
		http.Error(w, "Some error occured!!", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/profile", http.StatusFound)
}
