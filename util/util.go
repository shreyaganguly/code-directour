package util

import (
	"net/url"

	"github.com/unrolled/render"
)

var Renderer *render.Render
var Endpoint string

func SetRenderer(r *render.Render) {
	Renderer = r
}

func SetEndpoint(e string) {
	Endpoint = e
}

func IsLink(u string) bool {
	parsed, err := url.ParseRequestURI(u)

	if err != nil {
		return false
	}

	if len(parsed.Scheme) == 0 || len(parsed.Host) == 0 {
		return false
	}
	return true
}
