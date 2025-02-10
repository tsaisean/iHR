package model

import "time"

/* Note: GORM's AutoMigrate will not work for rename, drop a column or incompatible column type changes. Please do manual migration when necessary.*/

// TODO: Add multi devices support.

type Account struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `gorm:"size:20;unique;not null" json:"username"`
	Password  string    `gorm:"size:80;not null" json:"password"`
	Email     string    `gorm:"size:320" json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
