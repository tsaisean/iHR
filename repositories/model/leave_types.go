package model

type LeaveType struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"size:255;not null"`
	Description string `gorm:"size:255;not null"`
	Active      bool   `gorm:"default:true"`
}
