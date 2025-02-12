package authenticate

import (
	"github.com/gin-gonic/gin"
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

	auth, err := NewAuth(h.jwtSecret, "local", "none", acc.ID, acc.Username)
	if err := h.authRepo.CreateAuth(auth); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, auth)
}
