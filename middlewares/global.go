package middlewares

import (
	"net/http"

	interpose "github.com/carbocation/interpose/middleware"
	"github.com/dre1080/recover"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

// HandlePanic handle panics from your API requests.
func HandlePanic(handler http.Handler) http.Handler {
	recovery := recover.New(&recover.Options{
		Log: logrus.Print,
	})
	return recovery(handler)
}

// LogViaLogrus using interpose to integrate with logrus
func LogViaLogrus(handler http.Handler) http.Handler {
	logViaLogrus := interpose.NegroniLogrus()
	return logViaLogrus(handler)
}

// Cross creates a new http.Handler that adds authentication logic to a given Handler
func Cross(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Debugln("Enter into cross handler")
		logrus.Debugln("request: ", r)

		if r.Method == "OPTIONS" {
			if r.Header.Get("Access-Control-Request-Method") != "" {
				logrus.Debugf("cors preflight detected")
				// cors preflight request/response
				w.Header().Add("Access-Control-Allow-Origin", "*")
				w.Header().Add("Access-Control-Allow-Methods", "*")
				w.Header().Add("Access-Control-Allow-Headers", "*")
				w.Header().Add("Access-Control-Max-Age", "86400")
				w.Header().Add("Content-Type", "text/html; charset=utf-8")
				w.WriteHeader(200)

				if flusher, ok := w.(http.Flusher); ok {
					flusher.Flush()
				}
				return
			}
		}

		logrus.Debugln("Writing to header in callBefore \"Access-Control-Allow-Origin: *\"")

		w.Header().Add("Access-Control-Allow-Origin", "*")

		h.ServeHTTP(w, r)
	})
}

// RedocUI ...
func RedocUI(handler http.Handler) http.Handler {
	// return http.FileServer(assetFS())
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		opts := middleware.RedocOpts{
			SpecURL: r.URL.Host + "/dashboard/static/swagger/swagger.yaml",
			Title:   "User Manager",
		}
		middleware.Redoc(opts, handler).ServeHTTP(w, r)
		return
	})
}
