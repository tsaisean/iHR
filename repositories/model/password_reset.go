package model

import "time"

type PasswordReset struct {
	ID        uint   `gorm:"primary_key"`
	AccountID uint   `gorm:"index"`
	Token     string `gorm:"size:100;unique"`
	ExpiresAt time.Time
	Used      bool `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
