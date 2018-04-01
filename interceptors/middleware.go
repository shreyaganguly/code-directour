package interceptors

import (
	"net/http"

	"github.com/shreyaganguly/code-directour/util"
)

func AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userName := util.GetUserName(r)
		if userName == "" {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		next(w, r)
	}

}
