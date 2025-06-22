package db

import "time"

type User struct {
	ID          string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string    `gorm:"not null"`
	Email       string    `gorm:"uniqueIndex;not null"`
	Password    string    `gorm:"not null"`
	UTMSource   string    `gorm:""`
	UTMMedium   string    `gorm:""`
	UTMCampaign string    `gorm:""`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	LastLoginAt time.Time `gorm:"autoUpdateTime"`
	IsActive    bool      `gorm:"default:true"`
	IsAdmin     bool      `gorm:"default:false"`
}

type RequestLog struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Method    string
	Path      string
	IP        string
	UserAgent string
	Duration  int64          // in ms
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	Meta      map[string]any `gorm:"type:jsonb"`
}
