package util

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/shreyaganguly/code-directour/models"
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

func GenerateSharedTo(shareInfos []*models.ShareInfo) string {
	var s []string
	for _, shareinfo := range shareInfos {
		s = append(s, fmt.Sprintf(" %s (via %s)", shareinfo.SharedTo, shareinfo.Method))
	}
	reStr := regexp.MustCompile(",([^,]*)$")
	repStr := " and $1"
	return reStr.ReplaceAllString(strings.Join(s, ","), repStr)
}
