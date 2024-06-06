package router

import (
	"github.com/an-halim/golang-advanced/session-1-introduction-gin/handler"
	"github.com/an-halim/golang-advanced/session-1-introduction-gin/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", handler.GetHandler)

	privateRouter := r.Group("/private")
	privateRouter.Use(middleware.AuthMiddleware())
	{
		privateRouter.POST("/post", handler.PostHander)
	}
}
