package main

import (
	"context"
	"log"

	"github.com/an-halim/golang-advanced/session-7-pg-gorm/handler"
	"github.com/an-halim/golang-advanced/session-7-pg-gorm/repository/postgres_gorm"
	"github.com/an-halim/golang-advanced/session-7-pg-gorm/repository/postgres_pgx"
	"github.com/an-halim/golang-advanced/session-7-pg-gorm/router"
	"github.com/an-halim/golang-advanced/session-7-pg-gorm/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// setup gorm connectoin
	dsn := "postgresql://postgres:root@localhost:5432/golang_advance"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}
	log.Print("Connected to database : ", gormDB.Name())

	// pgxPool, err := pgxpool.New(context.Background(), "postgresql://postgres:root@localhost:5432/golang_advance")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Print("Connected to database")

	// setup service
	// var mockUserDBInSlice []entity.User
	// userRepo := slice.NewUserRepository(mockUserDBInSlice)
	// pgx db is enabled. comment to disabled
	// userRepo := postgres_pgx.NewUserRepository(pgxPool)
	// uncomment to use postgres gorm
	userRepo := postgres_gorm.NewUserRepository(gormDB)
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
