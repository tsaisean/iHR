package repositories

import (
	"iHR/repositories/model"

	"gorm.io/gorm"
)

//go:generate mockery --all --output=./mocks
type ResetPasswordRepository interface {
	CreatePasswordReset(reset *model.PasswordReset) error
	FindPasswordResetByToken(token string) (*model.PasswordReset, error)
	UpdatePasswordReset(reset *model.PasswordReset) error
	FindByEmail(email string) (*model.Account, error)
}

type resetPasswordRepo struct {
	db *gorm.DB
}

func (r *resetPasswordRepo) CreatePasswordReset(reset *model.PasswordReset) error {
	return r.db.Create(reset).Error
}

func (r *resetPasswordRepo) FindPasswordResetByToken(token string) (*model.PasswordReset, error) {
	var reset model.PasswordReset
	if err := r.db.Where("token = ?", token).First(&reset).Error; err != nil {
		return nil, err
	}
	return &reset, nil
}

func (r *resetPasswordRepo) UpdatePasswordReset(reset *model.PasswordReset) error {
	return r.db.Save(reset).Error
}

func (r *resetPasswordRepo) FindByEmail(email string) (*model.Account, error) {
	var account model.Account
	if err := r.db.Where("email = ?", email).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func NewResetPasswordRepo(db *gorm.DB) ResetPasswordRepository {
	return &resetPasswordRepo{
		db: db,
	}
}
