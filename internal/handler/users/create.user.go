package users

import (
	"encoding/json"
	"net/http"

	"server/internal/db"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	UTMSource   string `json:"utm_source,omitempty"`
	UTMMedium   string `json:"utm_medium,omitempty"`
	UTMCampaign string `json:"utm_campaign,omitempty"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// common function within the package for error handling
func respondWithError(w http.ResponseWriter, status int, errMsg string, detail string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error:   errMsg,
		Message: detail,
	})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "InvalidRequest", "Invalid JSON input")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "HashError", "Failed to secure password")
		return
	}

	user := db.User{
		Name:        req.Name,
		Email:       req.Email,
		Password:    string(hashedPassword),
		UTMSource:   req.UTMSource,
		UTMMedium:   req.UTMMedium,
		UTMCampaign: req.UTMCampaign,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, "DBError", "Could not save user to database")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created",
		"user_id": user.ID,
	})
}
