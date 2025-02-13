package model

import "time"

type LeaveRequest struct {
	ID          uint   `gorm:"primary_key"`
	CreatorID   uint   `gorm:"index"`
	EmployeeID  uint   `gorm:"index"`
	ApproverID  uint   `gorm:"index"`
	Status      string `gorm:"size:15;index"` // "pending", "approved", "used", "canceled", "rejected"
	LeaveTypeID uint   `gorm:"index"`
	StartDate   time.Time
	EndDate     time.Time
	Duration    uint   // Minutes
	Reason      string `gorm:"size:50"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
