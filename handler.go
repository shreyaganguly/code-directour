package main

import (
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	snippets, err := all(getUserName(r))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderer.HTML(w, http.StatusOK, "index", reverse(snippets))
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	renderer.HTML(w, http.StatusOK, "new", Languages)
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
	http.Redirect(w, r, "/index", http.StatusFound)
}
