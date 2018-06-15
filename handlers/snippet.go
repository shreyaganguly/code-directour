package handlers

import (
	"net/http"
	"net/mail"
	"strings"
	"time"

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
	user, err := db.LookupinUser(util.GetUserName(r))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		SnippetInfos models.Snippets
		ErrorMessage string
		User         *models.User
	}{
		snippets.Own().Reverse(),
		"",
		user,
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
	snippets, err := db.All(util.GetUserName(r))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err := db.LookupinUser(util.GetUserName(r))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		SnippetInfos models.Snippets
		ErrorMessage string
		User         *models.User
	}{
		snippets.Own().Reverse(),
		"",
		user,
	}
	util.Renderer.HTML(w, http.StatusOK, "all", data)
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
		snippet.ModifiedAt = time.Now().Unix()
		err = db.Update(snippet)
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
		return
	}
	snippet = models.NewSnippet(util.GetUserName(r), r.FormValue("title"), r.FormValue("language"), r.FormValue("code"), r.FormValue("references"), false, "", false, nil)
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

func sharedToListHandler(w http.ResponseWriter, r *http.Request) {
	snippets, err := db.All(util.GetUserName(r))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.Renderer.HTML(w, http.StatusOK, "sharedlist", snippets.SharedTo().Reverse())
}

func shareHandler(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	recepient := strings.ToLower(r.PostFormValue("recepient"))
	userExists := db.UserExists(recepient)
	if !userExists {
		snippets, err := db.All(util.GetUserName(r))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user, err := db.LookupinUser(util.GetUserName(r))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			SnippetInfos models.Snippets
			ErrorMessage string
			User         *models.User
		}{
			snippets.Own().Reverse(),
			"This User does not have a code-directour account!!!",
			user,
		}
		util.Renderer.HTML(w, http.StatusOK, "all", data)
		return
	}
	snippet, err := db.FindAndUpdate(util.GetUserName(r), args["key"], &models.ShareInfo{SharedTo: recepient, Method: "code-directour"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sharedSnippet := models.NewSnippet(recepient, snippet.Title, snippet.Language, snippet.Code, snippet.References, true, util.GetUserName(r), false, nil)
	err = db.Update(sharedSnippet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/all", http.StatusFound)
}

func shareEmailHandler(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	user, err := db.LookupinUser(util.GetUserName(r))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	models.NewMailer(user.Email)
	snippet, err := db.Find(util.GetUserName(r), args["key"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	models.SmtpMailer.Receiver = mail.Address{
		Address: r.PostFormValue("email"),
		Name:    r.PostFormValue("name"),
	}
	models.SmtpMailer.Data = snippet

	err = models.SmtpMailer.SendMail()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = db.FindAndUpdate(util.GetUserName(r), args["key"], &models.ShareInfo{SharedTo: r.PostFormValue("email"), Method: "email"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/all", http.StatusFound)
}

func shareSlackHandler(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	snippet, err := db.Find(util.GetUserName(r), args["key"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = UploadSnippet(snippet, r.PostFormValue("name"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = db.FindAndUpdate(util.GetUserName(r), args["key"], &models.ShareInfo{SharedTo: r.PostFormValue("name"), Method: "slack"})
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
