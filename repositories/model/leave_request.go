package model

import "time"

type LeaveRequest struct {
	ID          uint      `gorm:"primary_key"`
	CreatorID   uint      `gorm:"index"`
	EmployeeID  uint      `gorm:"index" json:"employee_id"`
	ApproverID  *uint     `gorm:"index"`
	Approver    Employee  `gorm:"foreignKey:ApproverID;references:ID"`
	Status      string    `gorm:"size:15;index"` // "pending", "approved", "used", "canceled", "rejected"
	LeaveTypeID uint      `gorm:"index" json:"leave_type_id"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Duration    uint      `json:"duration"` // Total Minutes
	Hours       uint      `gorm:"-" json:"hours"`
	Minutes     uint      `gorm:"-" json:"minutes"`
	Reason      string    `gorm:"size:50"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
