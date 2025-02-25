package authenticate

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"iHR/repositories/model"
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

	var emp *model.Employee
	if emp, err = h.empRepo.GetEmployeeByAccID(claims.UserID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Employee record not found"})
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auth, err := NewAuth(h.jwtSecret, "local", "none", claims.UserID, claims.Username, emp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.authRepo.CreateAuth(auth); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, auth)
}
