package middleware

import "net/http"

var contentTypes = []string{
	"application/proto",
	"application/grpc",
	"application/json",
	"application/connect+json",
	"application/grpc+proto",
	"application/grpc-web+proto",
}

func containsHeader(header string, headers []string) bool {
	for _, h := range headers {
		if h == header {
			return true
		}
	}
	return false
}

func AttachConnectFallback(next http.Handler, fallback http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if containsHeader(r.Header.Get("Content-Type"), contentTypes) {
			next.ServeHTTP(w, r)
			return
		}

		fallback.ServeHTTP(w, r)
	})
}

func ServeConnect(path string, handler http.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle(path, handler)
	return mux
}

func AttachConnect(mux *http.ServeMux) func(path string, handler http.Handler) {
	return func(path string, handler http.Handler) {
		if mux == nil {
			mux = http.NewServeMux()
		}

		mux.Handle(path, handler)
	}
}
