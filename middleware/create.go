package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

// mw - передаем список мидлваров
func CreateMiddleware(mw ...Middleware) func(http.Handler) http.Handler {
	return func(hnd http.Handler) http.Handler {
		next := hnd
		for k := len(mw) - 1; k >= 0; k-- {
			next = mw[k](next)
		}
		return next
	}
}
