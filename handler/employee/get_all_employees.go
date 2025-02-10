package employee

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"iHR/repositories/db"
	"iHR/repositories/model"
	"net/http"
	"strconv"
)

var DefaultPageSize = "20"

func validatePaginationParams(cursorStr, pageStr string) error {
	if cursorStr == "" && pageStr == "" {
		return errors.New("missing cursor or page query parameter")
	}
	if cursorStr != "" && pageStr != "" {
		return errors.New("can only have either cursor or page query parameter")
	}
	return nil
}

func (h *EmployeeHandler) GetAllEmployees(c *gin.Context) {
	cursorStr, pageStr := c.Query("cursor"), c.Query("page")

	if err := validatePaginationParams(cursorStr, pageStr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", DefaultPageSize))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pageSize."})
		return
	}

	var employees []model.Employee
	var currentPage int

	if cursorStr != "" {
		employees, currentPage, err = h.getEmployeesByCursor(c, cursorStr, pageSize)
	} else {
		employees, currentPage, err = h.getEmployeesByPage(c, pageStr, pageSize)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employees."})
		return
	}

	total, err := h.repo.GetTotal()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	totalPages := (total + pageSize - 1) / pageSize

	var nextPageCursorResponse *int
	if nextPageCursor := getNextPageCursor(employees); nextPageCursor > -1 {
		nextPageCursorResponse = &nextPageCursor
	}

	c.JSON(http.StatusOK, gin.H{
		"employees":    employees,
		"cursor":       nextPageCursorResponse,
		"current_page": currentPage,
		"total_pages":  totalPages,
		"total":        total,
	})

}

func (h *EmployeeHandler) getEmployeesByCursor(ctx context.Context, cursorStr string, pageSize int) ([]model.Employee, int, error) {
	cursor, err := strconv.Atoi(cursorStr)
	if err != nil {
		return nil, 0, errors.New("invalid cursor")
	}
	employees, err := h.repo.GetAllEmployeesAfter(ctx, cursor, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return employees, getCurrentPage(cursor, pageSize), nil
}

func (h *EmployeeHandler) getEmployeesByPage(ctx context.Context, pageStr string, pageSize int) ([]model.Employee, int, error) {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		return nil, 0, errors.New("invalid page number")
	}
	offset := (page - 1) * pageSize
	employees, err := h.repo.GetAllEmployeesFrom(ctx, offset, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return employees, page, nil
}

func getCurrentPage(cursor int, pageSize int) int {
	var recordsBeforeCursor int64
	db.DB.Model(&model.Employee{}).Where("id <= ?", cursor).Count(&recordsBeforeCursor)

	return int(recordsBeforeCursor)/pageSize + 1
}

func getNextPageCursor(employees []model.Employee) int {
	if len(employees) == 0 {
		return -1
	}

	return int(employees[len(employees)-1].ID)
}
