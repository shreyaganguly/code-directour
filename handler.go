package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderer.HTML(w, http.StatusOK, "index", nil)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	renderer.HTML(w, http.StatusOK, "new", LanguageMaps)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("*********", r.Form)
}
