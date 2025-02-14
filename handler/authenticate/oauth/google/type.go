package google

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"iHR/config"
	repo "iHR/repositories"
)

type GoogleOAuthHandler struct {
	jwtSecret    string
	oauth2Config *oauth2.Config
	accRepo      repo.AccountRepository
	authRepo     repo.AuthRepository
	empRepo      repo.EmployeeRepository
}

func NewGoogleOAuthHandler(jwtSecret string, oauthConfig config.Google, authRepo repo.AuthRepository, accRepo repo.AccountRepository, empRepo repo.EmployeeRepository) *GoogleOAuthHandler {
	var googleOauthConfig = &oauth2.Config{
		ClientID:     oauthConfig.ClientID,
		ClientSecret: oauthConfig.ClientSecret,
		// TODO: Change the domain dynamically for production/staging env
		RedirectURL: "http://localhost:8080/auth/google/callback",
		Scopes:      []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:    google.Endpoint,
	}

	return &GoogleOAuthHandler{jwtSecret: jwtSecret, oauth2Config: googleOauthConfig, accRepo: accRepo, authRepo: authRepo, empRepo: empRepo}
}
