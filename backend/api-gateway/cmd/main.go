package main

import (
	"fmt"
	"time"

	authentication "github.com/GeuberLucas/Gofre/backend/api-gateway/internal/Authentication"
	"github.com/GeuberLucas/Gofre/backend/api-gateway/internal/observability"
	"github.com/GeuberLucas/Gofre/backend/api-gateway/internal/reverseproxy"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	rp := reverseproxy.NewProxyRoutes()

	r := gin.New()
	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Info().
			Int("status", c.Writer.Status()).
			Dur("latency", time.Since(start)).
			Str("client_ip", c.ClientIP()).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Msg("")
	})

	r.Use(gin.Recovery())
	r.Use(authentication.AuthMiddleware(rp))
	r.Use(observability.CorrelationIDMiddleware())

	r.Any("/api/*path", rp.ProxyHandler)

	err := r.Run(":80")
	if err != nil {
		fmt.Printf("Api stopped with error: %v\n", err)
	}

}
