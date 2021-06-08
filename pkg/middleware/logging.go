package middleware

import (
	"fmt"
	"github.com/btassone/crypto/omega/pkg/web"
	"net/http"
	"net/http/httputil"
)

var (
	LoggingMiddleware = web.Middleware{
		Name: "logging",
		Func: Logging,
	}
)

// Logging Handles the general logging for the api routes
func Logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(requestDump))

		// Onto the next handler...
		next.ServeHTTP(w, r)
	}
}
