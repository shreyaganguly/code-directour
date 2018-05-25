package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shreyaganguly/code-directour/util"
)

func profileHandler(w http.ResponseWriter, r *http.Request) {
	util.Renderer.HTML(w, http.StatusOK, "profile", nil)
}

func profileSaveHandler(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	switch t := args["type"]; t {
	case "link":
		util.SetEndpoint(r.PostFormValue("endpoint"))
	case "email":
	case "slack":
	default:
		http.Error(w, "Some error occured!!", http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/profile", http.StatusFound)
}
