package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Paulinhoh/golang-basico/aula_05/01_proto-basic/pb"
	"google.golang.org/grpc"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *GreeterServer) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{
		Message: fmt.Sprintf("Olá, %s! Bem-vindo ao mundo dos gRPCs!", req.Name),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	fmt.Println("Servidor gRPC rodando na porta 50051...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Erro ao execultar o servidor %v", err)
	}
}
