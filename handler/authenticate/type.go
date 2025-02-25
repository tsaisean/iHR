package authenticate

import (
	repo "iHR/repositories"
	"iHR/services/email"
)

type AuthenticateHandler struct {
	jwtSecret         string
	accRepo           repo.AccountRepository
	authRepo          repo.AuthRepository
	empRepo           repo.EmployeeRepository
	resetPasswordRepo repo.ResetPasswordRepository
	emailService      *email.EmailService
}

func NewAuthenticateHandler(jwtSecret string, 
	accRepo repo.AccountRepository, 
	auth repo.AuthRepository, 
	empRepo repo.EmployeeRepository, 
	resetPasswordRepo repo.ResetPasswordRepository,
	emailService *email.EmailService) *AuthenticateHandler {
	return &AuthenticateHandler{jwtSecret: jwtSecret, accRepo: accRepo, authRepo: auth, empRepo: empRepo, resetPasswordRepo: resetPasswordRepo, emailService: emailService}
}
