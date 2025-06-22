package users

import (
	"encoding/json"
	"net/http"

	"server/internal/db"
	"server/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "InvalidRequest", "Invalid JSON input")
		return
	}

	// Find user by email
	var user db.User
	if err := db.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		respondWithError(w, http.StatusUnauthorized, "AuthFailed", "Invalid email or password")
		return
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		respondWithError(w, http.StatusUnauthorized, "AuthFailed", "Invalid email or password")
		return
	}

	generatedJWT, _ := utils.GenerateJWT(user.ID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{
		Message: "Login successful",
		Token:   generatedJWT,
	})
}
