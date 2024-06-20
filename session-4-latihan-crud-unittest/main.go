package main

import (
	"fmt"

	"github.com/an-halim/golang-advanced/session-4-latihan-crud/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	router.SetupRouter(r)

	fmt.Println("Server is running at :8000")
	r.Run(":8000")
}
