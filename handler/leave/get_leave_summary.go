package leave

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (l *LeaveHandler) GetLeaveSummary(c *gin.Context) {
	employeeIdStr := c.Query("employee_id")
	employeeId, err := strconv.Atoi(employeeIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	ls, err := l.repo.GetLeaveSummaries(uint(employeeId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, ls)
}
