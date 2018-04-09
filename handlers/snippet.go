package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shreyaganguly/code-directour/db"
	"github.com/shreyaganguly/code-directour/models"
	"github.com/shreyaganguly/code-directour/util"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	util.Renderer.HTML(w, http.StatusOK, "index", util.GetUserName(r))
}

func snippetsHandler(w http.ResponseWriter, r *http.Request) {
	snippets, err := db.All(util.GetUserName(r))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		SnippetInfos models.Snippets
		ErrorMessage string
		Endpoint     string
	}{
		snippets.Own().Reverse(),
		"",
		util.Endpoint,
	}
	util.Renderer.HTML(w, http.StatusOK, "all", data)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	util.Renderer.HTML(w, http.StatusOK, "new", models.Languages)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	snippet, err := db.Find(util.GetUserName(r), args["key"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		Languages []*models.Language
		Snippet   *models.SnippetInfo
	}{
		models.Languages,
		snippet,
	}
	util.Renderer.HTML(w, http.StatusOK, "edit", data)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {

	args := mux.Vars(r)
	err := db.Delete(util.GetUserName(r), args["key"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/all", http.StatusFound)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	var snippet *models.SnippetInfo
	var err error
	err = r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.FormValue("key") != "" {
		snippet, err = db.Find(util.GetUserName(r), r.FormValue("key"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = db.Delete(util.GetUserName(r), r.FormValue("key"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		snippet.Title = r.FormValue("title")
		snippet.Language = r.FormValue("language")
		snippet.Code = r.FormValue("code")
		snippet.References = r.FormValue("references")
	} else {
		snippet = models.NewSnippet(util.GetUserName(r), r.FormValue("title"), r.FormValue("language"), r.FormValue("code"), r.FormValue("references"), false, "", false, "")
	}
	err = db.Update(snippet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/all", http.StatusFound)
}

func shareListHandler(w http.ResponseWriter, r *http.Request) {
	snippets, err := db.All(util.GetUserName(r))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.Renderer.HTML(w, http.StatusOK, "sharedlist", snippets.Others().Reverse())
}

func shareHandler(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	recepient := r.PostFormValue("recepient")
	userExists := db.UserExists(recepient)
	if !userExists {
		snippets, err := db.All(util.GetUserName(r))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			SnippetInfos models.Snippets
			ErrorMessage string
			Endpoint     string
		}{
			snippets.Own().Reverse(),
			"This User does not have a code-directour account!!!",
			util.Endpoint,
		}
		util.Renderer.HTML(w, http.StatusOK, "all", data)
		return
	}
	snippet, err := db.FindAndUpdate(util.GetUserName(r), args["key"], recepient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sharedSnippet := models.NewSnippet(recepient, snippet.Title, snippet.Language, snippet.Code, snippet.References, true, util.GetUserName(r), false, "")
	err = db.Update(sharedSnippet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/all", http.StatusFound)
}

func shareEmailHandler(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	snippet, err := db.Find(util.GetUserName(r), args["key"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	models.SmtpMailer.Receiver.Address = r.PostFormValue("email")
	models.SmtpMailer.Data = snippet
	err = models.SmtpMailer.SendMail()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/all", http.StatusFound)
}

func linkHandler(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	snippet, err := db.Find(args["name"], args["key"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.Renderer.HTML(w, http.StatusOK, "link", snippet)
}
