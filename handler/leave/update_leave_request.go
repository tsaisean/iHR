package leave

import (
	"github.com/gin-gonic/gin"
	"iHR/repositories/model"
	"net/http"
	"strconv"
)

func (l *LeaveHandler) UpdateLeaveRequest(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	leaveRequest := &model.LeaveRequest{}
	if err := c.ShouldBindJSON(leaveRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if leaveRequest, err = l.repo.UpdateLeaveRequest(uint(id), leaveRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, leaveRequest)
}
