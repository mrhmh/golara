package controllers

import (
	"context"
	"golara/app/grpc/pb"
	"golara/core/facades/auth"
	"log"
)

type questionController struct {
}

func (s *questionController) Index(ctx context.Context, r *pb.PaginateRequest) (*pb.IndexResponse, error) {

	log.Println("Received: Question Index")

	userId := int32(auth.User().ID)
	return &pb.IndexResponse{Message: "Hello " + r.GetPerPage(), Code: userId}, nil
}

func NewQuestionController() *questionController {
	return &questionController{}
}
