package main

import (
	"context"
	"log"
	"net"

	pb "github.com/an-halim/golang-advanced/session-8-introduction-grpc/proto"
	userpb "github.com/an-halim/golang-advanced/session-8-introduction-grpc/proto/user"
	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

type UserServer struct {
	userpb.UnimplementedUserServiceServer
}

type UserEntity struct {
	Id    int32
	Name  string
	Email string
}

var Users []UserEntity
var lastId int32

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *UserServer) AddUser(ctx context.Context, in *userpb.User) (*userpb.Response, error) {
	lastId++

	User := UserEntity{Id: lastId, Name: in.Name, Email: in.Email}
	Users = append(Users, User)
	return &userpb.Response{
		Result: &userpb.Response_Create{
			Create: &userpb.ResponseCreate{
				Message: "User created successfully",
				Data:    &userpb.User{Id: User.Id, Name: User.Name, Email: User.Email},
			},
		},
	}, nil
}

func (s *UserServer) GetUser(ctx context.Context, in *userpb.ID) (*userpb.Response, error) {
	for _, user := range Users {
		if user.Id == in.Id {
			return &userpb.Response{
				Result: &userpb.Response_Create{
					Create: &userpb.ResponseCreate{
						Message: "User created successfully",
						Data:    &userpb.User{Id: user.Id, Name: user.Name, Email: user.Email},
					},
				},
			}, nil
		}
	}
	return nil, nil
}

func (s *UserServer) GetAllUser(ctx context.Context, in *userpb.Empty) (*userpb.Response, error) {
	var users []*userpb.User
	for _, user := range Users {
		users = append(users, &userpb.User{Id: user.Id, Name: user.Name, Email: user.Email})
	}
	return &userpb.Response{
		Result: &userpb.Response_GetAll{
			GetAll: &userpb.ResponseGetAll{
				Message: "Fetched all users successfully",
				Data:    users,
				Limit:   1,
				Page:    1,
			},
		},
	}, nil
}

func main() {
	runServer()
}

func runServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	userpb.RegisterUserServiceServer(s, &UserServer{})
	log.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
