package repositories

import (
	"gorm.io/gorm"
	"iHR/db/model"
)

//go:generate mockery --all --output=./mocks
type AuthRepository interface {
	CreateAuth(auth *model.Auth) error
	GetAuth(userID uint) (*model.Auth, error)
	InvalidateAuth(userID uint) error
}

type AuthRepo struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) CreateAuth(auth *model.Auth) error {
	// Invalidate all the user's auth before create a new one.
	if err := r.InvalidateAuth(auth.UserID); err != nil {
		return err
	}
	return r.db.Create(auth).Error
}

func (r *AuthRepo) GetAuth(userID uint) (*model.Auth, error) {
	var auth model.Auth
	if err := r.db.First(&auth, userID).Error; err != nil {
		return nil, err
	}
	return &auth, nil
}

func (r *AuthRepo) InvalidateAuth(userID uint) error {
	return r.db.Delete(&model.Auth{}, "user_id = ?", userID).Error
}
