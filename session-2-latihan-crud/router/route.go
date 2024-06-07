package router

import (
	"github.com/an-halim/golang-advanced/session-2-latihan-crud/handler"
	"github.com/an-halim/golang-advanced/session-2-latihan-crud/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	usersPublicEndpoint := r.Group("/users")
	usersPublicEndpoint.GET("/", handler.GetUsers)
	usersPublicEndpoint.GET("/:id", handler.GetUserbyID)

	usersPrivateEndpoint := r.Group("/users")
	usersPrivateEndpoint.Use(middleware.AuthMiddleware())
	usersPrivateEndpoint.POST("/", handler.CreateUser)
	usersPrivateEndpoint.PUT("/:id", handler.UpdateUser)
	usersPrivateEndpoint.DELETE("/:id", handler.DeleteUser)
}
