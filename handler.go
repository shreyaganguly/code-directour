package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderer.HTML(w, http.StatusOK, "index", getUserName(r))
}

func snippetsHandler(w http.ResponseWriter, r *http.Request) {
	snippets, err := all(getUserName(r))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderer.HTML(w, http.StatusOK, "all", reverse(snippets))
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	renderer.HTML(w, http.StatusOK, "new", Languages)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	snippet, err := findSnippetForUser(getUserName(r), args["key"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		Languages []*Language
		Snippet   *SnippetInfo
	}{
		Languages,
		snippet,
	}
	renderer.HTML(w, http.StatusOK, "edit", data)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	snippet := NewSnippet(r.FormValue("title"), r.FormValue("language"), r.FormValue("code"), r.FormValue("references"))
	err = snippet.Save(getUserName(r))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/all", http.StatusFound)
}
