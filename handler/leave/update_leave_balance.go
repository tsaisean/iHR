package leave

import (
	"github.com/gin-gonic/gin"
	"iHR/repositories/model"
	"net/http"
	"strconv"
)

func (l *LeaveHandler) UpdateLeaveBalance(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	leaveBalance := &model.LeaveBalances{}
	if err := c.ShouldBindJSON(leaveBalance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if leaveBalance, err = l.repo.UpdateLeaveBalance(uint(id), leaveBalance); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, leaveBalance)
}
