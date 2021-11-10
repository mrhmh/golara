package grpc

import (
	"golara/app/grpc/controllers"
	"golara/app/grpc/pb"
	"google.golang.org/grpc"
)

func RegisterRoutes(s *grpc.Server) {

	pb.RegisterQuestionServer(s, controllers.NewQuestionController())

}
