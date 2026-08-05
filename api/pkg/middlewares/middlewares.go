package middlewares

import (
	"log/slog"
	"net/http"
	"os"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		logger.Info("\n host:%s requestpath: %s  method:%s  remoteIp:%s", r.Host, r.RequestURI, r.Method, r.RemoteAddr)
		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		next(w, r)
	}
}
