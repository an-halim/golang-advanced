package main

import (
	"context"
	"log"
	"net"
	"time"

	grpcHandler "github.com/an-halim/golang-advanced/session-9-crud-user-grpc-gateway/handler/grpc"
	pb "github.com/an-halim/golang-advanced/session-9-crud-user-grpc-gateway/proto/user_service/v1"
	"github.com/an-halim/golang-advanced/session-9-crud-user-grpc-gateway/repository/postgres_gorm"
	"github.com/an-halim/golang-advanced/session-9-crud-user-grpc-gateway/service"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	grpcServers "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	log.Print("Connected to database : ", gormDB.Config)

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

	go func() {
		log.Println("Running grpc server in port :50051")
		_ = grpcServer.Serve(lis)
	}()
	time.Sleep(1 * time.Second)

	// run grpc gateway
	conn, err := grpc.NewClient("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gwmux := runtime.NewServeMux()

	if err := pb.RegisterUserServiceHandler(context.Background(), gwmux, conn); err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}

	gwServer := gin.Default()

	gwServer.Group("/v1/*{grpc_gateway}").Any("", gin.WrapH(gwmux))
	log.Println("Running grpc gateway in port :8000")
	_ = gwServer.Run(":8000")

}
