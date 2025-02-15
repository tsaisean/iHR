package employee

import (
	"github.com/gin-gonic/gin"
	. "iHR/handler/employee/response"
	"iHR/repositories"
	"net/http"
)

func (h *EmployeeHandler) AutoComplete(c *gin.Context) {
	query := c.Query("query")

	if len(query) < 3 {
		c.JSON(http.StatusOK, AutoCompleteResponse{Suggestions: []repositories.Suggestion{}})
		return
	}

	if suggestions, err := h.repo.Autocomplete(c, query); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	} else {
		response := AutoCompleteResponse{
			Suggestions: suggestions,
		}
		c.JSON(http.StatusOK, response)
	}
}
