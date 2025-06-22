package users

import (
	"encoding/json"
	"net/http"

	"server/internal/middleware"
)

func GetMe(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(string)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Token is valid",
		"user_id": userID,
	})
}
