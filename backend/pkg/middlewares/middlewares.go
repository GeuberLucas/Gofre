package middlewares

import (
	"net/http"

	jwtToken "github.com/GeuberLucas/Gofre/backend/pkg/jwt"
	"github.com/GeuberLucas/Gofre/backend/pkg/response"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc{
	return func( w http.ResponseWriter, r *http.Request){
		if err := jwtToken.ValidateToken(r); err != nil{
			 response.ErrorResponse(w,http.StatusUnauthorized,err)
			 return
		}
		next(w,r)
	}
}

