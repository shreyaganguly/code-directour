package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shreyaganguly/code-directour/db"
	"github.com/shreyaganguly/code-directour/models"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderer.HTML(w, http.StatusOK, "index", getUserName(r))
}

func snippetsHandler(w http.ResponseWriter, r *http.Request) {
	snippets, err := db.All(getUserName(r))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		SnippetInfos models.Snippets
		ErrorMessage string
		Endpoint     string
	}{
		reverse(snippets.Own()),
		"",
		*endpoint,
	}
	renderer.HTML(w, http.StatusOK, "all", data)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	renderer.HTML(w, http.StatusOK, "new", models.Languages)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	snippet, err := db.Find(getUserName(r), args["key"])
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
	renderer.HTML(w, http.StatusOK, "edit", data)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {

	args := mux.Vars(r)
	err := db.Delete(getUserName(r), args["key"])
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
		snippet, err = db.Find(getUserName(r), r.FormValue("key"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = db.Delete(getUserName(r), r.FormValue("key"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		snippet.Title = r.FormValue("title")
		snippet.Language = r.FormValue("language")
		snippet.Code = r.FormValue("code")
		snippet.References = r.FormValue("references")
	} else {
		snippet = models.NewSnippet(getUserName(r), r.FormValue("title"), r.FormValue("language"), r.FormValue("code"), r.FormValue("references"), false, "", false, "")
	}
	err = db.Update(snippet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/all", http.StatusFound)
}

func shareListHandler(w http.ResponseWriter, r *http.Request) {
	snippets, err := db.All(getUserName(r))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderer.HTML(w, http.StatusOK, "sharedlist", reverse(snippets.Others()))
}

func shareHandler(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	recepient := r.PostFormValue("recepient")
	userExists := db.UserExists(recepient)
	if !userExists {
		snippets, err := db.All(getUserName(r))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			SnippetInfos models.Snippets
			ErrorMessage string
			Endpoint     string
		}{
			reverse(snippets.Own()),
			"This User does not have a code-directour account!!!",
			*endpoint,
		}
		renderer.HTML(w, http.StatusOK, "all", data)
		return
	}
	snippet, err := db.FindAndUpdate(getUserName(r), args["key"], recepient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sharedSnippet := models.NewSnippet(recepient, snippet.Title, snippet.Language, snippet.Code, snippet.References, true, getUserName(r), false, "")
	err = db.Update(sharedSnippet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/all", http.StatusFound)
}

func linkHandler(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	snippet, err := db.Find(getUserName(r), args["key"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderer.HTML(w, http.StatusOK, "link", snippet)
}
