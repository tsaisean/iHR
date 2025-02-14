package model

import "time"

/* Note: GORM's AutoMigrate will not work for rename, drop a column or incompatible column type changes. Please do manual migration when necessary.*/

type Employee struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	AccID        uint      `gorm:"index" json:"account_id"`
	FirstName    string    `gorm:"size:50" json:"first_name"`
	LastName     string    `gorm:"size:50" json:"last_name"`
	Email        string    `gorm:"size:320" json:"email"`
	Position     string    `gorm:"size:50" json:"position"`
	SupervisorID *uint     `gorm:"index" json:"supervisor_id"`
	Salary       uint      `gorm:"default:0" json:"salary"`
	CreatedAt    time.Time `json:"created_at"`
}
