package model

import "time"

type Auth struct {
	ID                    uint      `gorm:"primary_key;auto_increment" json:"id"`
	UserID                uint      `gorm:"index" json:"user_id"`
	Token                 string    `gorm:"-" json:"token"`
	RefreshToken          string    `gorm:"size:350" json:"refresh_token"`
	CreatedAt             time.Time `json:"created_at"`
	TokenExpiresAt        time.Time `json:"token_expires_at"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
}
