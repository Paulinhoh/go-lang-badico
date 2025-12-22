package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Paulinhoh/golang-basico/aula_05/02_api-project/handler"
	"github.com/Paulinhoh/golang-basico/aula_05/02_api-project/middleware"
	pb "github.com/Paulinhoh/golang-basico/aula_05/02_api-project/pb"
	"github.com/Paulinhoh/golang-basico/aula_05/02_api-project/service"
	"google.golang.org/grpc"
)

func main() {
	// inicializar o serviço de usuarios userService :=
	userService := service.NewUserService()

	// configurar servidor gRPC
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor gRPC: %v", err)
	}

	// configurar interceptor para o servidor gRPC
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			middleware.UnaryLoggingInterceptor,
		),
	}
	grpcServer := grpc.NewServer(serverOptions...)
	userServiceServer := handler.NewUserServiceServer(userService)

	pb.RegisterUserServiceServer(
		grpcServer,
		userServiceServer,
	)
	fmt.Println("Servidor gRPC rodando na porta 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Erro ao execultar o servidor gRPC: %v", err)
	}
}
