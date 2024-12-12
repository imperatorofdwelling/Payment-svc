package middleware

import (
	"github.com/google/uuid"
	"net/http"
)

func IdempotenceKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost || r.Method == http.MethodDelete {
			key := uuid.New()
			r.Header.Set("Idempotence-Key", key.String())
		}

		next.ServeHTTP(w, r)
	})
}
