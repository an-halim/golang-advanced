package main

import (
	"context"
	"log"

	"github.com/an-halim/golang-advanced/session-6-pg-database/handler"
	"github.com/an-halim/golang-advanced/session-6-pg-database/repository/postgres_pgx"
	"github.com/an-halim/golang-advanced/session-6-pg-database/router"
	"github.com/an-halim/golang-advanced/session-6-pg-database/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	pgxPool, err := connectDB("postgresql://postgres:root@localhost:5432/golang_advance")
	if err != nil {
		log.Fatalln(err)
	}
	log.Print("Connected to database")

	// setup service
	// var mockUserDBInSlice []entity.User
	// userRepo := slice.NewUserRepository(mockUserDBInSlice)
	// pgx db is enabled. comment to disabled
	userRepo := postgres_pgx.NewUserRepository(pgxPool)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Routes
	router.SetupRouter(r, userHandler)

	log.Print("Server started at :8000")
	r.Run(":8000")
}

func connectDB(dbURL string) (postgres_pgx.PgxPoolIface, error) {
	return pgxpool.New(context.Background(), dbURL)
}
