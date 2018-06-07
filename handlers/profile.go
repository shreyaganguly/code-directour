package handlers

import (
	"fmt"
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
	}
	switch t := args["type"]; t {
	case "link":
		r.ParseForm()
		fmt.Println(r.Form)
		user.Endpoint = r.PostFormValue("endpoint")
		err = db.Update(user)
		if err != nil {
			//TODO: cleaner error handling
			http.Error(w, "Some error occured!!", http.StatusInternalServerError)
			return
		}
	case "email":
		user.Email = models.NewEmailSettings(r.PostFormValue("server"), r.PostFormValue("port"), r.PostFormValue("email"), r.PostFormValue("password"), r.PostFormValue("sendername"), r.PostFormValue("senderemail"))
		if user.Email.Server != "" && user.Email.Port != "" && user.Email.Address != "" && user.Email.Password != "" {
			user.Email.Enabled = true
		} else {
			user.Email.Enabled = false
		}
		err = db.Update(user)
		if err != nil {
			//TODO: cleaner error handling
			http.Error(w, "Some error occured!!", http.StatusInternalServerError)
			return
		}
	case "slack":
		user.Slack = &models.Slack{Token: r.PostFormValue("token")}
		err = db.Update(user)
		if err != nil {
			//TODO: cleaner error handling
			http.Error(w, "Some error occured!!", http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "Some error occured!!", http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/profile", http.StatusFound)
}
