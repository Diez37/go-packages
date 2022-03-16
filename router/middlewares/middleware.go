package middlewares

import "net/http"

type Middleware interface {
	Middleware(next http.Handler) http.Handler
}

// MiddlewareHandler general type of middleware with golang standard interfaces
type MiddlewareHandler func(next http.Handler) http.Handler
