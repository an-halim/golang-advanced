package main

import (
	"fmt"

	"github.com/an-halim/golang-advanced/session-5-validator/entity"
	"github.com/an-halim/golang-advanced/session-5-validator/handler"
	slice "github.com/an-halim/golang-advanced/session-5-validator/repository/silce"
	"github.com/an-halim/golang-advanced/session-5-validator/router"
	"github.com/an-halim/golang-advanced/session-5-validator/service"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// setup service
	var mockUserDBInSlice []entity.User
	userRepo := slice.NewUserRepository(mockUserDBInSlice)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router.SetupRouter(r, userHandler)

	fmt.Println("Server is running at :8000")
	r.Run(":8000")
}
