package google

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"iHR/config"
	"iHR/repositories"
)

type GoogleOAuthHandler struct {
	jwtSecret    string
	oauth2Config *oauth2.Config
	accRepo      repositories.AccountRepository
	authRepo     repositories.AuthRepository
}

func NewGoogleOAuthHandler(jwtSecret string, oauthConfig config.Google, authRepo repositories.AuthRepository, accRepo repositories.AccountRepository) *GoogleOAuthHandler {
	var googleOauthConfig = &oauth2.Config{
		ClientID:     oauthConfig.ClientID,
		ClientSecret: oauthConfig.ClientSecret,
		// TODO: Change the domain dynamically for production/staging env
		RedirectURL: "http://localhost:8080/auth/google/callback",
		Scopes:      []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:    google.Endpoint,
	}

	return &GoogleOAuthHandler{jwtSecret: jwtSecret, oauth2Config: googleOauthConfig, accRepo: accRepo, authRepo: authRepo}
}
