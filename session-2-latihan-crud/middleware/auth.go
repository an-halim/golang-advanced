package middleware

import (
	"net/http"

	"github.com/an-halim/golang-advanced/session-2-latihan-crud/response"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()

		if !ok {
			apiResponse := response.APIResponseFailed{}
			apiResponse.Status = "failed"
			apiResponse.Message = "Authorization required"

			c.JSON(http.StatusUnauthorized, apiResponse)
			c.Abort()
			return
		}

		const (
			expectedUsername = "user"
			expectedPassword = "pass"
		)

		isValid := (username == expectedUsername) && (password == expectedPassword)

		if !isValid {
			apiResponse := response.APIResponseFailed{}
			apiResponse.Status = "failed"
			apiResponse.Message = "Invalid username or password"

			c.JSON(http.StatusUnauthorized, apiResponse)
			c.Abort()
			return
		}

		c.Next()
	}

}
