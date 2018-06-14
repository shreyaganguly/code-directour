package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shreyaganguly/code-directour/db"
	"github.com/shreyaganguly/code-directour/models"
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
		return
	}
	switch t := args["type"]; t {
	case "link":
		user.Endpoint = r.PostFormValue("endpoint")
	case "email":
		user.Email = models.NewEmailSettings(r.PostFormValue("server"), r.PostFormValue("port"), r.PostFormValue("email"), r.PostFormValue("password"), r.PostFormValue("sendername"), r.PostFormValue("senderemail"))
		if user.Email.Server != "" && user.Email.Port != "" && user.Email.Address != "" && user.Email.Password != "" {
			user.Email.Enabled = true
		} else {
			user.Email.Enabled = false
		}
	case "slack":
		user.Slack = &models.Slack{Token: r.PostFormValue("token")}
	default:
		http.Error(w, "not matching with any routes", http.StatusInternalServerError)
		return
	}
	err = db.Update(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.Renderer.HTML(w, http.StatusOK, "profile", user)
}
