package repositories

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"iHR/repositories/model"
)

//go:generate mockery --all --output=./mocks
type AccountRepository interface {
	CreateAccount(*model.Account) error
	Authenticate(username, password string) (*model.Account, error)
	GetIDByGoogleID(id string) (uint, error)
}

type AccountRepo struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepo {
	return &AccountRepo{db: db}
}

func (r *AccountRepo) CreateAccount(a *model.Account) error {
	return r.db.Create(a).Error
}

func (r *AccountRepo) Authenticate(username, password string) (*model.Account, error) {
	var account model.Account
	if err := r.db.Where("username = ?", username).First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("invalid username or password")
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("invalid username or password") // Avoid revealing which field failed.
	}

	return &account, nil
}

func (r *AccountRepo) GetIDByGoogleID(id string) (uint, error) {
	var userID uint

	if err := r.db.Model(&model.Account{}).Select("id").Where("google_id = ?", id).First(&userID).Error; err != nil {
		return 0, err
	}

	return userID, nil
}
