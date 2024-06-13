package middleware

import (
	"net/http"

	"github.com/an-halim/golang-advanced/session-2-latihan-crud-unittest/config"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware adalah middleware untuk autentikasi
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verifikasi token (misalnya, cocokkan dengan token yang diharapkan)
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization basic token required"})
			c.Abort()
			return
		}

		isValid := (username == config.AuthBasicUsername) && (password == config.AuthBasicPassword)
		if !isValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		// Lanjutkan ke handler berikutnya jika token valid
		c.Next()
	}
}