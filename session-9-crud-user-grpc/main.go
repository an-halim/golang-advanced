package main

import (
	"log"
	"net"

	grpcHandler "github.com/an-halim/golang-advanced/session-9-crud-user-grpc/handler/grpc"
	pb "github.com/an-halim/golang-advanced/session-9-crud-user-grpc/proto/user_service/v1"
	"github.com/an-halim/golang-advanced/session-9-crud-user-grpc/repository/postgres_gorm"
	"github.com/an-halim/golang-advanced/session-9-crud-user-grpc/service"
	"github.com/gin-gonic/gin"
	grpcServers "google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	// r := gin.Default()

	// setup gorm connectoin
	dsn := "postgresql://postgres:root@localhost:5432/golang_advance"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}
	log.Print("Connected to database : ", gormDB.Name())

	userRepo := postgres_gorm.NewUserRepository(gormDB)
	userService := service.NewUserService(userRepo)
	// userHandler := handler.NewUserHandler(userService)
	userGrpc := grpcHandler.NewUserHandler(userService)

	// Routes
	// router.SetupRouter(r, userHandler)
	grpcServer := grpcServers.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userGrpc)
	// log.Print("Server started at :8000")
	// r.Run(":8000")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Running grpc server in port :50051")
	_ = grpcServer.Serve(lis)
}
