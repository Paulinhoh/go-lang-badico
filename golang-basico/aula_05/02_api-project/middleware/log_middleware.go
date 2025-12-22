package middleware

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// UnaryInterceptor para log
func UnaryLoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	// execulta o handler (processa a requisição)
	res, err := handler(ctx, req)
	// log de requisição
	log.Printf(
		"Method: %s | Time: %s | Error: %v",
		info.FullMethod,
		time.Since(start),
		status.Convert(err).Message(),
	)

	return res, err
}
