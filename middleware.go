package main

import "net/http"

func authenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userName := getUserName(r)
		if userName == "" {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		next(w, r)
	}

}
