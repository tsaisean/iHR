package leave

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (l *LeaveHandler) GetLeaveRequests(c *gin.Context) {
	employeeIDStr := c.Query("employee_id")
	if employeeID, err := strconv.Atoi(employeeIDStr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required parameter \"employee_id\""})
		return
	} else {
		if leaveRequests, err := l.repo.GetAllLeaveRequests(uint(employeeID)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, leaveRequests)
		}
		return
	}
}
