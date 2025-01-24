package authenticate

import repo "iHR/db/repositories"

type AuthenticateHandler struct {
	jwtSecret string
	accRepo   repo.AccountRepository
	authRepo  repo.AuthRepository
}

func NewAuthenticateHandler(jwtSecret string, accRepo repo.AccountRepository, auth repo.AuthRepository) *AuthenticateHandler {
	return &AuthenticateHandler{jwtSecret: jwtSecret, accRepo: accRepo, authRepo: auth}
}
