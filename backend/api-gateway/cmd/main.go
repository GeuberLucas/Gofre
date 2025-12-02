package main

import (
	"fmt"
	"net/http"

	"github.com/GeuberLucas/Gofre/backend/api-gateway/internal/proxyRoutes"
)

func main() {
	rp := proxyRoutes.NewProxyRoutes()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/", rp.ServeHTTP)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Api stopped with error: %v\n", err)
	}

}
