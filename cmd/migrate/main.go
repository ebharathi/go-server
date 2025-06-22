package main

import (
	"log"
	"os"
	"server/internal/db"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("❌ .env file not found")
		}
	}
	db.Init()
	log.Println("🔄 Starting migration...")

	if err := db.DB.AutoMigrate(
		&db.User{},
		&db.RequestLog{},
	); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("✅ Migration complete")
}
