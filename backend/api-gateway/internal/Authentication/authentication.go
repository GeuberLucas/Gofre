package authentication

import (
	"log"
	"net/http"
	"strings"

	"github.com/GeuberLucas/Gofre/backend/api-gateway/internal/reverseproxy"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(rp *reverseproxy.ReverseProxy) gin.HandlerFunc {
	return func(c *gin.Context) {

		path := c.Request.URL.Path

		isAuthRoute := strings.HasPrefix(path, "/api/auth")
		isProfile := strings.HasPrefix(path, "/api/auth/profile")

		// Rotas de auth não passam por verificação (login, register etc)
		if isAuthRoute && !isProfile {
			c.Next()
			return
		}

		userID, err := rp.VerifyAuthentication(c, c.Request.Header)
		if err != nil {
			log.Println("Auth error:", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		// Guarda no contexto do Gin
		c.Set("user_id", userID)

		c.Next()
	}
}
