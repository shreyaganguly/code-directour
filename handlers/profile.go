package handlers

import (
	"net/http"

	"github.com/shreyaganguly/code-directour/util"
)

func profileHandler(w http.ResponseWriter, r *http.Request) {
	util.Renderer.HTML(w, http.StatusOK, "profile", nil)
}
