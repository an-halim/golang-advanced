package main

import (
	"fmt"

	"github.com/an-halim/golang-advanced/session-4-latihan-crud-unittest/entity"
	"github.com/an-halim/golang-advanced/session-4-latihan-crud-unittest/handler"
	slice "github.com/an-halim/golang-advanced/session-4-latihan-crud-unittest/repository/silce"
	"github.com/an-halim/golang-advanced/session-4-latihan-crud-unittest/router"
	"github.com/an-halim/golang-advanced/session-4-latihan-crud-unittest/service"
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
