package middlewares

import (
	"log"
	"net/http"

	jwtToken "github.com/GeuberLucas/Gofre/backend/pkg/jwt"
	"github.com/GeuberLucas/Gofre/backend/pkg/response"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n host:%s requestpath: %s  method:%s  remoteIp:%s", r.Host, r.RequestURI, r.Method, r.RemoteAddr)
		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := jwtToken.ValidateToken(r); err != nil {
			response.ErrorResponse(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
