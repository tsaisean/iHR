package employee

import (
	"github.com/gin-gonic/gin"
	"iHR/repositories/model"
	"iHR/utils"
	"net/http"
)

func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	employee := new(model.Employee)
	if err := c.ShouldBindJSON(employee); err != nil {
		if isUnmarshalError, msg := utils.GetUnmarshalTypeErrorMsg(err); isUnmarshalError {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var err error
	if employee, err = h.repo.CreateEmployee(c, employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, employee)
}
