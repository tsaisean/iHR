package model

import "time"

type Auth struct {
	ID                    uint   `gorm:"primary_key;auto_increment"`
	UserID                uint   `gorm:"index"`
	Token                 string `gorm:"size:350"`
	RefreshToken          string `gorm:"size:350"`
	CreatedAt             time.Time
	TokenExpiresAt        time.Time
	RefreshTokenExpiresAt time.Time
}
