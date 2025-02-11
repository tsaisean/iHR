package authenticate

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"iHR/repositories/model"
	"iHR/utils"
	"net/http"
	"regexp"
)

type signupForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"` // For verification
}

var errorMessages = map[string]string{
	"Username": "Username is required.",
	"Password": "Password is required.",
	"Email":    "A valid email address is required.",
}

func validatePassword(password string) error {
	// Length check (8 to 16 characters)
	if len(password) < 8 || len(password) > 16 {
		return errors.New("password must be between 8 and 16 characters long")
	}

	// Check for at least one lowercase letter
	lowercase := regexp.MustCompile(`[a-z]`)
	if !lowercase.MatchString(password) {
		return errors.New("password must contain at least one lowercase letter")
	}

	// Check for at least one uppercase letter
	uppercase := regexp.MustCompile(`[A-Z]`)
	if !uppercase.MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}

	// Check for at least one digit
	digit := regexp.MustCompile(`\d`)
	if !digit.MatchString(password) {
		return errors.New("password must contain at least one digit")
	}

	// Check for at least one special character (@$!%*?&)
	specialChar := regexp.MustCompile(`[@$!%*?&]`)
	if !specialChar.MatchString(password) {
		return errors.New("password must contain at least one special character (@$!%*?&)")
	}

	return nil // Password is valid
}

func (h *AuthenticateHandler) Signup(c *gin.Context) {
	var form signupForm
	if err := c.ShouldBindJSON(&form); err != nil {
		if isUnmarshalError, msg := utils.GetUnmarshalTypeErrorMsg(err); isUnmarshalError {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": msg})
			return
		} else if isMissingFieldError, msg := utils.GetMissingFieldErrorMsg(err, errorMessages); isMissingFieldError {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": msg})
			return
		} else if validateErr := validatePassword(form.Password); validateErr != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": validateErr.Error()})
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
