package authenticate

import (
	repo "iHR/repositories"
)

type AuthenticateHandler struct {
	jwtSecret string
	accRepo   repo.AccountRepository
	authRepo  repo.AuthRepository
	empRepo   repo.EmployeeRepository
	resetPasswordRepo repo.ResetPasswordRepository
}

func NewAuthenticateHandler(jwtSecret string, accRepo repo.AccountRepository, auth repo.AuthRepository, empRepo repo.EmployeeRepository, resetPasswordRepo repo.ResetPasswordRepository) *AuthenticateHandler {
	return &AuthenticateHandler{jwtSecret: jwtSecret, accRepo: accRepo, authRepo: auth, empRepo: empRepo, resetPasswordRepo: resetPasswordRepo}
}
