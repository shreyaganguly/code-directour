package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/shreyaganguly/code-directour/db"
	"github.com/shreyaganguly/code-directour/handlers"
	"github.com/shreyaganguly/code-directour/models"
	"github.com/shreyaganguly/code-directour/util"
	"github.com/unrolled/render"
)

var (
	dbPath   = flag.String("db", "directour.db", "File to store the db")
	host     = flag.String("b", "0.0.0.0", "Host to start your code-directeur")
	port     = flag.Int("p", 8080, "Port to start your code-directeur")
	endpoint = flag.String("e", "http://0.0.0.0:8080", "Endpoint that will be shared in the link")
)

var (
	renderer *render.Render
	location *time.Location
)

func main() {
	//TODO : change view structure
	//TODO: add recently deleted snippet section
	//TODO: add sharing history
	//TODO: add show more / less in view
	//TODO: remove views as html
	// TODO: add mailgun and slackbot
	// TODO: give link from listing of snippets to a particular snippet
	// TODO: give share action in overflow button
	//TODO: make edit and delete post request
	//TODO: refactor code
	// TODO: add date created/modified in the listing section
	//TODO: change name of functions and methods and bucketnames
	// TODO: add comments for exported functions
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *host, *port)
	viewHelpers := template.FuncMap{
		"getCode":    models.GetCode,
		"getAceCode": models.GetAceCode,
		"IsLink":     util.IsLink,
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
	util.SetEndpoint(*endpoint)
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
