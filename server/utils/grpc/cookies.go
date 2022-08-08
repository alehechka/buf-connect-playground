package grpc

import (
	"net/http"
	"strings"

	connect_go "github.com/bufbuild/connect-go"
)

// AddCookie adds the defined cookie has a Set-Cookie header in the response
func AddCookie[T any](res *connect_go.Response[T], cookie *http.Cookie) {
	res.Header().Add("Set-Cookie", cookie.String())
}

// GetCookie locates and returns the value of the cookie with matching name
func GetCookie[T any](req *connect_go.Request[T], name string) (value string) {
	cookies := req.Header().Values("Cookie")

	for _, cookie := range cookies {
		for _, cookieName := range strings.Split(cookie, ";") {
			trimmed := strings.Trim(cookieName, " ")
			if strings.HasPrefix(trimmed, name) {
				return trimmed[len(name)+1:]
			}
		}
	}

	return
}
