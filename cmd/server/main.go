package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"server/internal/db"
	"server/internal/router"
	"server/internal/utils"
)

func main() {
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("❌ .env file not found")
		}
	}

	utils.InitJWT()
	utils.InitGoogleOAuth()
	db.Init()

	r := router.SetupRouter() // Setup routes
	log.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r)) // Start HTTP server
}
