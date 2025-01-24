package authenticate

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"iHR/db/model"
	"iHR/utils"
	"net/http"
)

func (h *AuthenticateHandler) Signup(c *gin.Context) {
	type signupForm struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email" binding:"required"` // For verification
	}

	var form signupForm
	if err := c.ShouldBindJSON(&form); err != nil {
		if isUnmarshalError, msg := utils.GetUnmarshalTypeErrorMsg(err); isUnmarshalError {
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	acc := model.Account{
		Username: form.Username,
		Password: string(hashedPassword),
	}

	if err := h.accRepo.CreateAccount(&acc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully!"})
}
