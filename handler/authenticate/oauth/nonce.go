package oauth

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
)
import "github.com/gin-contrib/sessions"

func generateRandomString() string {
	bytes := make([]byte, 16) // 16 bytes = 128-bit secure random value
	_, err := rand.Read(bytes)
	if err != nil {
		return "" // Handle error properly in production
	}
	return base64.URLEncoding.EncodeToString(bytes)
}

func GenerateNonce(c *gin.Context) string {
	nonce := generateRandomString()
	StoreNonceInSession(c, nonce)
	return nonce
}

func StoreNonceInSession(c *gin.Context, nonce string) {
	session := sessions.Default(c)
	session.Set("oauth_nonce", nonce)
	session.Save()
}

// Validate nonce from session
func IsValidNonceFromSession(c *gin.Context, nonce string) bool {
	session := sessions.Default(c)
	storedNonce := session.Get("oauth_nonce")

	if storedNonce == nil || storedNonce != nonce {
		return false
	}

	// Remove nonce after validation
	session.Delete("oauth_nonce")
	session.Save()

	return true
}
