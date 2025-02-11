package employee

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	. "iHR/repositories/model"
	"net/http"
	"strconv"
)

func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee := new(Employee)
	if err = c.ShouldBindJSON(employee); err != nil {
		var unmarshalTypeError *json.UnmarshalTypeError
		if errors.As(err, &unmarshalTypeError) {
			message := fmt.Sprintf("Field '%s' has a type mismatch. Expected '%s'.", unmarshalTypeError.Field, unmarshalTypeError.Type)
			c.JSON(http.StatusBadRequest, gin.H{"error": message})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload."})
		return
	}

	employee, err = h.repo.UpdateEmployeeByID(c, uint(id), employee)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "record not found."})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, employee)
}
