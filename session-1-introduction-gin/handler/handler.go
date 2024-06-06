package handler

import (
	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Halo dari gin!",
	})
}

func PostHander(c *gin.Context) {
	var json struct {
		Message string `json:"message"`
	}

	if err := c.ShouldBindJSON(&json); err == nil {
		c.JSON(200, gin.H{"message": json.Message})
	} else {
		c.JSON(400, gin.H{"error": err.Error()})
	}
}
