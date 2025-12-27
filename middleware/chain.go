package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func Chain(final http.Handler, middlewares ...Middleware) http.Handler {
	h := final

	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}

	return h
}
