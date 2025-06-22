package middleware

import (
	"net/http"
	"time"

	"server/internal/db"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start).Milliseconds()
		metadata := make(map[string]any)

		// Only add userID if it exists in context
		if userID, ok := r.Context().Value(UserIDKey).(string); ok {
			metadata["user_id"] = userID
		}

		db.DB.Create(&db.RequestLog{
			Method:    r.Method,
			Path:      r.URL.Path,
			IP:        r.RemoteAddr,
			UserAgent: r.UserAgent(),
			Duration:  duration,
			Meta:      metadata,
		})
	})
}
