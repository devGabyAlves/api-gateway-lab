package middleware

import (
	"log"
	"net/http"
	"strings"
)

func SanitizeForLog(input string) string {
	input = strings.ReplaceAll(input, "\n", " ")
	input = strings.ReplaceAll(input, "\r", " ")
	return input
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	    ip := SanitizeForLog(r.RemoteAddr)
	    method := SanitizeForLog(r.Method)
	    path := SanitizeForLog(r.URL.Path)

		log.Printf(
			"[REQ] ip=%s method=%s path=%s",
			ip,
			method,
			path,
		)
		next.ServeHTTP(w, r)
	})
}


