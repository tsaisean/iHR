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
