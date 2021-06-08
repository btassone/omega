package web

import (
	"net/http"
)

// MiddlewareFunc Middleware function signature
type MiddlewareFunc func(next http.HandlerFunc) http.HandlerFunc

// Middleware The base middleware struct
type Middleware struct {
	Name string
	Func MiddlewareFunc
}

// RegisterMiddleware Registers multiple middleware and returns the configured HandlerFunc
func RegisterMiddleware(h http.HandlerFunc, m []Middleware) http.HandlerFunc {
	if len(m) < 1 {
		return h
	}

	wrapped := h

	// loop in reverse to preserve middleware order
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i].Func(wrapped)
	}

	return wrapped

}
