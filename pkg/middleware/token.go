package middleware

import (
	"github.com/btassone/crypto/omega/pkg/web"
	"net/http"
)

var (
	TokenMiddleware = web.Middleware{
		Name: "token",
		Func: Token,
	}
)

// Token Handles token authentication for the app private routes
func Token(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Token logic

		// Onto the next handler...
		next.ServeHTTP(w, r)
	}
}
