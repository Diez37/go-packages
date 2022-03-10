package middlewares

import "net/http"

// Middleware general type of middleware with golang standard interfaces
type Middleware func(next http.Handler) http.Handler
