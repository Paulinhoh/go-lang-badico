package handler

import (
	"context"
	"fmt"

	pb "github.com/Paulinhoh/golang-basico/aula_05/02_api-project/pb"
	"github.com/Paulinhoh/golang-basico/aula_05/02_api-project/service"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	service *service.UserService
}

func NewUserServiceServer(service *service.UserService) *UserServiceServer {
	return &UserServiceServer{service: service}
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	id, err := s.service.CreateUser(req.Name, int(req.Age))
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		Id:      id,
		Message: "usuario criado com sucesso!",
	}, nil
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.service.GetUserById(req.Id)
	if err != nil {
		return nil, fmt.Errorf("usuario não encontrado: %v", err)
	}

	return &pb.GetUserResponse{
		Id:   user.ID,
		Name: user.Name,
		Age:  int32(user.Age),
	}, nil
}
