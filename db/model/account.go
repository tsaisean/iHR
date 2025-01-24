package model

import "time"

/* Note: GORM's AutoMigrate will not work for rename, drop a column or incompatible column type changes. Please do manual migration when necessary.*/

// TODO: Add multi devices support.

type Account struct {
	ID        uint   `gorm:"primary_key"`
	Username  string `gorm:"size:20;unique;not null"`
	Password  string `gorm:"size:80;not null"`
	Email     string `gorm:"size:320"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
