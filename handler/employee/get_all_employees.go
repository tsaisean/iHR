package employee

import (
	"github.com/gin-gonic/gin"
	"iHR/db"
	"iHR/db/model"
	"net/http"
	"strconv"
)

var DefaultPageSize = "20"

func (h *EmployeeHandler) GetAllEmployees(c *gin.Context) {
	cursorStr := c.Query("cursor")
	pageStr := c.Query("page")

	if cursorStr == "" && pageStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing cursor or page query parameter."})
		return
	} else if cursorStr != "" && pageStr != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can only have either one cursor or page query parameter."})
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
		cursor, err := strconv.Atoi(cursorStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cursor."})
			return
		}
		employees, err = h.repo.GetAllEmployeesAfter(cursor, pageSize)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		currentPage = getCurrentPage(cursor, pageSize)
	} else {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page."})
			return
		}
		if page < 1 {
			page = 1
		}
		offset := (page - 1) * pageSize
		employees, err = h.repo.GetAllEmployeesFrom(offset, pageSize)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		currentPage = page
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
