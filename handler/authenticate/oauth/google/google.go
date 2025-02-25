package google

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	_ "golang.org/x/oauth2"
	"gorm.io/gorm"
	"iHR/handler/authenticate"
	"iHR/handler/authenticate/oauth"
	"iHR/repositories/model"
	"log"
	"net/http"
)

func generateOAuthState(flowType string, nonce string) string {
	stateData := map[string]string{
		"flow":  flowType, // Can be "signup" or "login"
		"nonce": nonce,    // Helps with CSRF protection
	}

	// Convert stateData to JSON
	stateJSON, _ := json.Marshal(stateData)

	// Base64 encode the JSON string (so it can be safely passed in URL)
	return base64.URLEncoding.EncodeToString(stateJSON)
}

func (g *GoogleOAuthHandler) Signup(c *gin.Context) {
	nonce := oauth.GenerateNonce(c)
	url := g.oauth2Config.AuthCodeURL(generateOAuthState("signup", nonce))
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (g *GoogleOAuthHandler) Login(c *gin.Context) {
	nonce := oauth.GenerateNonce(c)
	url := g.oauth2Config.AuthCodeURL(generateOAuthState("login", nonce))
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (g *GoogleOAuthHandler) Callback(c *gin.Context) {
	stateParam := c.Query("state")

	stateBytes, err := base64.URLEncoding.DecodeString(stateParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state"})
		return
	}

	var stateData map[string]string
	if err := json.Unmarshal(stateBytes, &stateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state JSON"})
		return
	}

	flowType, nonce := stateData["flow"], stateData["nonce"]
	if !oauth.IsValidNonceFromSession(c, nonce) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid CSRF token"})
		return
	}

	code := c.Query("code")
	token, err := g.oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		log.Println("Failed to exchange code for token", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token."})
		return
	}

	client := g.oauth2Config.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Println("Failed to get user info.", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		log.Println("Failed to decode user info", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode user info"})
		return
	}

	acc := model.Account{
		GoogleID: userInfo["id"].(*string),
	}

	if flowType == "signup" {
		err := g.accRepo.CreateAccount(&acc)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create account."})
		}
		return
	} else {
		acc.ID, err = g.accRepo.GetIDByGoogleID(*acc.GoogleID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found."})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
			}
			return
		}
	}

	var emp *model.Employee
	if emp, err = g.empRepo.GetEmployeeByAccID(acc.ID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Employee record not found"})
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auth, err := authenticate.NewAuth(g.jwtSecret, "oauth", "google", acc.ID, userInfo["name"].(string), emp)
	if err := g.authRepo.CreateAuth(auth); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, auth)
}
