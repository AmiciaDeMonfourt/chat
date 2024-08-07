package middleware

import (
	"log/slog"
	"net/http"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request:", "method", r.Method, "url", r.URL)
		next.ServeHTTP(w, r)
	})
}
