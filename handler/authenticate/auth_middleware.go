package authenticate

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

func (h *AuthenticateHandler) AuthMiddleware(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format."})
		c.Abort()
		return
	}

	if _, err := ValidateToken(h.jwtSecret, tokenString); err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is expired"})
			c.Abort()
			return
		}

		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}
