package util

import "github.com/unrolled/render"

var Renderer *render.Render
var Endpoint string

func SetRenderer(r *render.Render) {
	Renderer = r
}

func SetEndpoint(e string) {
	Endpoint = e
}
