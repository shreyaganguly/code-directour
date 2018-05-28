package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/mail"
	"time"

	"github.com/shreyaganguly/code-directour/db"
	"github.com/shreyaganguly/code-directour/handlers"
	"github.com/shreyaganguly/code-directour/models"
	"github.com/shreyaganguly/code-directour/util"
	"github.com/unrolled/render"
)

var (
	dbPath          = flag.String("db", "directour.db", "File to store the db")
	host            = flag.String("b", "0.0.0.0", "Host to start your code-directour")
	port            = flag.Int("p", 8080, "Port to start your code-directour")
	smtpServer      = flag.String("s", "smtp.gmail.com", "Host name of the SMTP Server")
	smtpPort        = flag.Int("t", 587, "SMTP port")
	smptpUser       = flag.String("u", "", "Username for SMTP authentication (If not passed sharing code snippets through email is disabled)")
	smtpPassword    = flag.String("w", "", "Password for SMTP authentication (If not passed sharing code snippets through email is disabled)")
	mailSenderName  = flag.String("sendername", "Code Directour", "Sender name")
	mailSenderEmail = flag.String("sendermail", "no-reply@code-directour.com", "Sender email")
)

var (
	renderer *render.Render
	location *time.Location
)

func main() {
	//TODO : change view structure
	//TODO: add recently deleted snippet section
	//TODO: add show more / less in view
	//TODO: remove views as html
	// TODO: give link from listing of snippets to a particular snippet
	// TODO: give share action in overflow button
	//TODO: make edit and delete post request
	//TODO: refactor code
	//TODO: change name of functions and methods and bucketnames
	// TODO: add comments for exported functions
	//TODO: add validations
	//TODO: add name along with email while sharing content through mail
	// TODO: show error message or success messages
	//TODO: handle slack errors more elegantly
	//TODO: all errors to be passed by ajax call
	//TODO: add more languages
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *host, *port)
	viewHelpers := template.FuncMap{
		"getCode":        models.GetCode,
		"getAceCode":     models.GetAceCode,
		"IsLink":         util.IsLink,
		"SharedtoString": util.GenerateSharedTo,
		"GetEndpoint":    util.GetEndpoint,
	}
	renderer = render.New(render.Options{
		Directory:       "views",
		Layout:          "layout",
		Extensions:      []string{".tmpl", ".html"},
		Funcs:           []template.FuncMap{viewHelpers},
		IsDevelopment:   true,
		RequirePartials: true,
	})
	util.SetRenderer(renderer)
	models.NewMailer(*smtpServer, *smtpPort, mail.Address{Name: *mailSenderName, Address: *mailSenderEmail}, mail.Address{}, *smptpUser, *smtpPassword, nil)
	err := db.Init(*dbPath)
	if err != nil {
		log.Fatal("Problem in initializing db  ", err)
	}
	location, err = time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Fatal("Problem in loadLocation  ", err)
	}
	models.SetLocation(location)
	log.Println("Starting code-directour at ", addr)
	err = http.ListenAndServe(addr, handlers.SetUpRoutes())
	if err != nil {
		log.Fatal(err)
	}
}
