package model

import "time"

type LeaveBalances struct {
	ID          uint64 `gorm:"primary_key" json:"id"`
	EmployeeID  uint   `gorm:"index" json:"employee_id"`
	LeaveTypeID uint   `gorm:"index" json:"leave_type_id"`
	Allocated   uint   `json:"allocated"` // Minutes
	StartDate   time.Time
	Status      string `gorm:"size:12;index" json:"status"`
}
