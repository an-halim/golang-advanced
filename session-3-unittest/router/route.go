package router

import (
	"github.com/an-halim/golang-advanced/session-3-unittest/handler"
	"github.com/an-halim/golang-advanced/session-3-unittest/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", handler.GetHandler)

	privateRouter := r.Group("/private")
	privateRouter.Use(middleware.AuthMiddleware())
	{
		privateRouter.POST("/post", handler.PostHandler)
	}
}
