package middleware

import (
	"github.com/btassone/crypto/omega/pkg/web"
	"net/http"
)

var (
	CorsMiddleware = web.Middleware{
		Name: "cors",
		Func: Cors,
	}
)

// Cors Handles setting up the CORS headers
func Cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Headers", "X-Auth-Token")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte{})
			return
		}

		// Onto the next handler...
		next.ServeHTTP(w, r)
	}
}
