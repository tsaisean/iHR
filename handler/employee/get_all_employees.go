package employee

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"iHR/db"
	"iHR/db/model"
	"log"
	"net/http"
	"strconv"
	"time"
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

	pageSizeStr := c.DefaultQuery("pageSize", DefaultPageSize)
	pageSize, err := strconv.Atoi(pageSizeStr)
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
		employees, err = h.getAllEmployeesByCursor(c, cursor, pageSize)
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
		employees, err = h.getAllEmployeesByPage(c, page, pageSize)
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

func (h *EmployeeHandler) getAllEmployeesByCursor(c context.Context, cursor int, pageSize int) ([]model.Employee, error) {
	cacheKey, employees, err := h.getFromCache(c, cursor, 0, pageSize)
	if err != nil {
		return nil, err
	} else if employees != nil {
		return employees, nil
	}

	employees, err = h.repo.GetAllEmployeesAfter(cursor, pageSize)
	if err != nil {
		return nil, err
	}

	h.setCache(c, cacheKey, employees)

	return employees, nil
}

func (h *EmployeeHandler) getAllEmployeesByPage(c context.Context, page int, pageSize int) ([]model.Employee, error) {
	cacheKey, employees, err := h.getFromCache(c, 0, page, pageSize)
	if err != nil {
		return nil, err
	} else if employees != nil {
		return employees, nil
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	employees, err = h.repo.GetAllEmployeesFrom(offset, pageSize)
	if err != nil {
		return nil, err
	}

	h.setCache(c, cacheKey, employees)

	return employees, nil
}

func (h *EmployeeHandler) getFromCache(c context.Context, cursor int, page int, pageSize int) (string, []model.Employee, error) {
	cacheKey := getGetCacheKey(cursor, page, pageSize)
	cache, err := h.cache.Get(c, cacheKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return cacheKey, nil, nil
		} else {
			return cacheKey, nil, err
		}
	}

	var employees []model.Employee
	err = json.Unmarshal([]byte(cache), &employees)
	if err != nil {
		return cacheKey, nil, err
	}

	log.Printf("Hit the cache for key: %s.", cacheKey)
	return cacheKey, employees, nil
}

func getGetCacheKey(cursor int, page int, pageSize int) string {
	return fmt.Sprintf("employees_c:%d_p:%d_ps:%d", cursor, page, pageSize)
}

func (h *EmployeeHandler) setCache(c context.Context, key string, employees []model.Employee) {
	if data, err := json.Marshal(employees); err != nil {
		log.Println("Error marshalling employees!")
	} else {
		h.cache.Set(c, key, data, 5*time.Minute)
	}
}
