package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/unrolled/render"
)

var (
	host = flag.String("b", "0.0.0.0", "Host to start your code-directeur")
	port = flag.Int("p", 8080, "Port to start your code-directeur")
)

var (
	renderer *render.Render
)

func main() {
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *host, *port)
	renderer = render.New(render.Options{
		Directory:       "views",
		Layout:          "layout",
		Extensions:      []string{".tmpl", ".html"},
		IsDevelopment:   true,
		RequirePartials: true,
	})
	err := http.ListenAndServe(addr, setupRoutes())
	if err != nil {
		log.Fatal(err)
	}
}
