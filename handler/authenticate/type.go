package authenticate

import (
	"iHR/repositories"
)

type AuthenticateHandler struct {
	jwtSecret string
	accRepo   repositories.AccountRepository
	authRepo  repositories.AuthRepository
}

func NewAuthenticateHandler(jwtSecret string, accRepo repositories.AccountRepository, auth repositories.AuthRepository) *AuthenticateHandler {
	return &AuthenticateHandler{jwtSecret: jwtSecret, accRepo: accRepo, authRepo: auth}
}
