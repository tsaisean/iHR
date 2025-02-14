package repositories

import (
	"gorm.io/gorm"
	. "iHR/repositories/model"
)

//go:generate mockery --all --output=./mocks
type LeaveRepository interface {
	CreateLeaveRequest(request *LeaveRequest) (*LeaveRequest, error)
	UpdateLeaveRequest(id uint, updated *LeaveRequest) (*LeaveRequest, error)
	GetAllLeaveRequests(employeeID uint) ([]LeaveRequest, error)
	CreateLeaveBalance(balance *LeaveBalances) (*LeaveBalances, error)
	UpdateLeaveBalance(id uint, updated *LeaveBalances) (*LeaveBalances, error)
}

type LeaveRepo struct {
	db *gorm.DB
}

func NewLeaveRepo(db *gorm.DB) *LeaveRepo {
	return &LeaveRepo{db: db}
}

func (r *LeaveRepo) CreateLeaveRequest(request *LeaveRequest) (*LeaveRequest, error) {
	if err := r.db.Create(&request).Error; err != nil {
		return nil, err
	}
	return request, nil
}

func (r *LeaveRepo) UpdateLeaveRequest(id uint, updated *LeaveRequest) (*LeaveRequest, error) {
	request := &LeaveRequest{}
	if err := r.db.First(request, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(request).Updates(updated).Error; err != nil {
		return nil, err
	}

	return request, nil
}

func (r *LeaveRepo) GetAllLeaveRequests(employeeID uint) ([]LeaveRequest, error) {
	var leaveRequests []LeaveRequest
	approverFieldsFunc := func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "first_name", "last_name")
	}

	if err := r.db.Preload("Approver", approverFieldsFunc).Where("employee_id = ?", employeeID).Find(&leaveRequests).Error; err != nil {
		return nil, err
	}

	return leaveRequests, nil
}

func (r *LeaveRepo) CreateLeaveBalance(balance *LeaveBalances) (*LeaveBalances, error) {
	if err := r.db.Create(balance).Error; err != nil {
		return nil, err
	}
	return balance, nil
}

func (r *LeaveRepo) UpdateLeaveBalance(id uint, updated *LeaveBalances) (*LeaveBalances, error) {
	balance := &LeaveBalances{}
	if err := r.db.First(balance, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(balance).Updates(updated).Error; err != nil {
		return nil, err
	}

	return balance, nil
}

type LeaveSummary struct {
	LeaveTypeID       uint   `json:"leave_type_id"`
	LeaveType         string `json:"leave_type"`
	Total             uint   `json:"total"`
	Used              uint   `json:"used"`
	PendingAllocation uint   `json:"pending_allocation"`
	Available         uint   `json:"available"`
}

func (r *LeaveRepo) GetLeaveSummaries(employeeID uint) ([]LeaveSummary, error) {
	var results []LeaveSummary

	err := r.db.
		Table("leave_types AS lt").
		Select(`
            lt.id AS leave_type_id,
            lt.name AS leave_type,
            
            COALESCE(SUM(lb.allocated), 0) AS total,

            COALESCE(
                SUM(
                    CASE WHEN lr.status IN ('pending','approved','used')
                    THEN lr.duration ELSE 0 END
                ), 
            0) AS used,

            COALESCE(
                SUM(
                    CASE WHEN lb.status = 'pending' 
                    THEN lb.allocated ELSE 0 END
                ),
            0) AS pending_allocation,

            COALESCE(
                SUM(lb.allocated)
                - SUM(
                    CASE WHEN lr.status IN ('pending','approved','used')
                    THEN lr.duration ELSE 0 END
                )
                - SUM(
                    CASE WHEN lr.status = 'approved'
                    THEN lr.duration ELSE 0 END
                ),
            0) AS available
        `).
		Joins(`
            LEFT JOIN leave_balances AS lb
            ON lb.leave_type_id = lt.id
            AND lb.employee_id = ?
        `, employeeID).
		Joins(`
            LEFT JOIN leave_requests AS lr
            ON lr.leave_type_id = lt.id
            AND lr.employee_id = ?
        `, employeeID).
		Group("lt.id, lt.name").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}
	return results, nil
}
