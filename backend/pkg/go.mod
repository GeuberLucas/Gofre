module github.com/GeuberLucas/Gofre/backend/pkg

go 1.25.4

require github.com/lib/pq v1.10.9

require github.com/gorilla/mux v1.8.1

require github.com/GeuberLucas/Gofre/backend/middlewares v0.0.0-20251112235549-9db0970dca8a

replace github.com/GeuberLucas/Gofre/backend/middlewares => ../middlewares
