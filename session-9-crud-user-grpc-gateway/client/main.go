package main

import (
	"context"

	"log"

	pb "github.com/an-halim/golang-advanced/session-9-crud-user-grpc-gateway/proto/user_service/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	runClient()
}

func runClient() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	userClient := pb.NewUserServiceClient(conn)

	name := "world"
	r, err := userClient.CreateUser(context.Background(), &pb.CreateUserRequest{Name: name, Email: "email", Password: "password"})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Printf("User: %s", r.GetMessage())

}
