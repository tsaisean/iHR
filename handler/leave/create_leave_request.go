package leave

import (
	"github.com/gin-gonic/gin"
	"iHR/repositories/model"
	"net/http"
	"time"
)

func (l *LeaveHandler) CreateLeaveRequest(c *gin.Context) {
	request := &model.LeaveRequest{}

	var employee *model.Employee
	if emp, exist := c.Get("employee"); !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "missing employeeID"})
		return
	} else {
		employee = emp.(*model.Employee)
	}

	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
	request.Duration = request.Hours*60 + request.Minutes
	request.ApproverID = employee.SupervisorID
	request.Status = "pending"

	var err error
	if request, err = l.repo.CreateLeaveRequest(request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, request)
}
