package grpc

import (
	v1controller "golara/app/grpc/controllers/v1"
	v1server "golara/app/grpc/pb/v1"
	"google.golang.org/grpc"
)

func RegisterRoutes(s *grpc.Server) {

	v1server.RegisterChallengeServer(s, v1controller.NewChallengeController())

}
