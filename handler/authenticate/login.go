package authenticate

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"iHR/repositories/model"
	"iHR/utils"
	"net/http"
)

func (h *AuthenticateHandler) Login(c *gin.Context) {
	type loginForm = struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	form := loginForm{}

	if err := c.ShouldBindJSON(&form); err != nil {
		if isUnmarshalError, msg := utils.GetUnmarshalTypeErrorMsg(err); isUnmarshalError {
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	acc, err := h.accRepo.Authenticate(form.Username, form.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var emp *model.Employee
	if emp, err = h.empRepo.GetEmployeeByAccID(acc.ID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Employee record not found"})
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auth, err := NewAuth(h.jwtSecret, "local", "none", acc.ID, acc.Username, emp)
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
