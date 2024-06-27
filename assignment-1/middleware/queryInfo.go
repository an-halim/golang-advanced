package middleware

import (
	"net/http"
	"strconv"

	"github.com/an-halim/golang-advanced/assignment-1/response"
	"github.com/gin-gonic/gin"
)

// PaginationMiddleware validates the 'size' and 'page' query parameters
func PaginationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		pageSize := c.DefaultQuery("size", "10")
		currentPage := c.DefaultQuery("page", "1")

		pageSizeInt, err := strconv.Atoi(pageSize)
		if err != nil || pageSizeInt <= 0 {

			apiResponse := response.Failed{}
			apiResponse.Status = "failed"
			apiResponse.Message = "Invalid page size"
			c.JSON(http.StatusBadRequest, apiResponse)
			c.Abort()
			return
		}

		// Convert currentPage to integer
		currentPageInt, err := strconv.Atoi(currentPage)
		if err != nil || currentPageInt <= 0 {
			apiResponse := response.Failed{}
			apiResponse.Status = "failed"
			apiResponse.Message = "Invalid current page"
			c.JSON(http.StatusBadRequest, apiResponse)
			c.Abort()
			return
		}

		offset := pageSizeInt * (currentPageInt - 1)

		c.Set("pageSize", pageSizeInt)
		c.Set("currentPage", currentPageInt)
		c.Set("offset", offset)

		c.Next()
	}
}
