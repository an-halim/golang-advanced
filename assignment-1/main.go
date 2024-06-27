package main

import (
	"log"

	"github.com/an-halim/golang-advanced/assignment-1/entity"
	"github.com/an-halim/golang-advanced/assignment-1/handler"
	postgres_gorm "github.com/an-halim/golang-advanced/assignment-1/repository"
	"github.com/an-halim/golang-advanced/assignment-1/router"
	"github.com/an-halim/golang-advanced/assignment-1/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	VALIDATOR := validator.New()

	// setup gorm connectoin
	dsn := "postgresql://postgres:root@localhost:5432/assignment_1"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}
	log.Print("Connected to database : ", gormDB.Name())
	log.Print("Starting Migrations")
	gormDB.AutoMigrate(&entity.User{}, &entity.Submission{})

	userRepo := postgres_gorm.NewUserRepository(gormDB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService, VALIDATOR)

	submissionRepo := postgres_gorm.NewSubmissionRepository(gormDB)
	submissionService := service.NewSubmissionService(submissionRepo, userService)
	submissionHandler := handler.NewSubmissionHandler(submissionService, VALIDATOR)

	// Routes
	router.SetupRouter(r, userHandler, submissionHandler)

	log.Print("Server started at :8000")
	r.Run(":8000")
}
