package authenticate

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

	auth, err := NewAuth(h.jwtSecret, claims.UserID, claims.Username)
	if err := h.authRepo.CreateAuth(auth); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, auth)
}

func ValidateToken(secret string, tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
