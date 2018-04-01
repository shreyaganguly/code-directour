package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

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
	// TODO: add sharing history
	//TODO: add show more / less in view
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *host, *port)
	viewHelpers := template.FuncMap{
		"getCode":    getCode,
		"getAceCode": getAceCode,
	}
	renderer = render.New(render.Options{
		Directory:       "views",
		Layout:          "layout",
		Extensions:      []string{".tmpl", ".html"},
		Funcs:           []template.FuncMap{viewHelpers},
		IsDevelopment:   true,
		RequirePartials: true,
	})
	err := initDB(*dbPath)
	if err != nil {
		log.Fatal("Problem in initializing db  ", err)
	}
	location, err = time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Fatal("Problem in loadLocation  ", err)
	}
	log.Println("Starting code-directour at ", addr)
	err = http.ListenAndServe(addr, setupRoutes())
	if err != nil {
		log.Fatal(err)
	}
}
