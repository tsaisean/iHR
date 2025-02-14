package leave

import (
	"github.com/gin-gonic/gin"
	"iHR/repositories/model"
	"net/http"
)

func (l *LeaveHandler) CreateLeaveBalance(c *gin.Context) {
	balance := &model.LeaveBalances{}

	if err := c.ShouldBindJSON(balance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var err error
	if balance, err = l.repo.CreateLeaveBalance(balance); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, balance)
}
