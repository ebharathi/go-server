package oauth

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"server/internal/db"
	"server/internal/utils"
)

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing code", http.StatusBadRequest)
		return
	}

	if utils.GoogleOAuthConfig == nil {
		http.Error(w, "Google OAuth not configured", http.StatusInternalServerError)
		return
	}

	token, err := utils.GoogleOAuthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Token exchange failed", http.StatusUnauthorized)
		return
	}

	client := utils.GoogleOAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var userInfo struct {
		Email string `json:"email"`
	}

	json.Unmarshal(body, &userInfo)

	var user db.User
	err = db.DB.Where("email = ?", userInfo.Email).First(&user).Error
	if err != nil {
		log.Println("New User Signed in via google")
		user = db.User{
			Email:     userInfo.Email,
			Name:      "",
			CreatedAt: time.Now(),
		}
		db.DB.Create(&user)
	}

	jwt, _ := utils.GenerateJWT(user.ID)

	// ✅ Set token in a secure cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    jwt,
		Path:     "/",
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	frontendURL := os.Getenv("FRONTEND_URL")

	http.Redirect(w, r, frontendURL, http.StatusFound)
}
