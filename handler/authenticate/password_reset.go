package authenticate

import (
	"crypto/rand"
	"encoding/hex"
	"iHR/repositories/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *AuthenticateHandler) RequestPasswordReset(c *gin.Context) {
	var request struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	account, err := h.resetPasswordRepo.FindByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "If the email exists, a reset link will be sent"})
		return
	}

	// Generate token
	token := generateResetToken()

	// Create password reset record
	reset := &model.PasswordReset{
		AccountID: account.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	if err := h.resetPasswordRepo.CreatePasswordReset(reset); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process request"})
		return
	}

	if err := h.emailService.SendPasswordResetEmail(account.Username, account.Email, token); err != nil {
		log.Printf("Failed to send reset email: %v", err)
		// Don't return error to client to prevent email enumeration
	}

	c.JSON(http.StatusOK, gin.H{"message": "If the email exists, a reset link will be sent"})
}

func (h *AuthenticateHandler) ResetPassword(c *gin.Context) {
	var request struct {
		Token    string `json:"token" binding:"required"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	reset, err := h.resetPasswordRepo.FindPasswordResetByToken(request.Token)
	if err != nil || reset.Used || time.Now().After(reset.ExpiresAt) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired token"})
		return
	}

	// Update password
	if err := h.accRepo.UpdatePassword(reset.AccountID, request.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	// Mark token as used
	reset.Used = true
	if err := h.resetPasswordRepo.UpdatePasswordReset(reset); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update reset status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password successfully reset"})
}

func generateResetToken() string {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}
