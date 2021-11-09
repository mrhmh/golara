package providers

import (
	"fmt"
	appGrpc "golara/app/grpc"
	"golara/app/grpc/interceptors"
	"golara/core/facades"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type GrpcServerProvider struct{}

func (p *GrpcServerProvider) Boot() {

	port := facades.Config().App.GrpcPORT

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	authInterceptor := new(interceptors.AuthInterceptor)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.Unary()),
	)
	reflection.Register(s)

	//register servers
	appGrpc.RegisterServers(s)

	fmt.Printf("Listening and serving gRPC on :%d", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
