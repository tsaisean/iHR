package authenticate

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *AuthenticateHandler) RefreshToken(c *gin.Context) {
	refreshToken := c.GetHeader("Refresh-Token")
	var claims *Claims
	var err error

	if claims, err = ValidateToken(h.jwtSecret, refreshToken); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token."})
		return
	}

	auth, err := NewAuth(h.jwtSecret, "local", "none", claims.UserID, claims.Username)
	if err := h.authRepo.CreateAuth(auth); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, auth)
}
