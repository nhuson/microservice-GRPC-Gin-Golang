package handler

import (
	"context"
	"log"
	pb "micr-go/services/users/pb"
)

type Users struct {
}

func (u *Users) FindAll(ctx context.Context, req *pb.FindAllRequest) (*pb.FindAllResponse, error) {
	log.Printf("Receive message %s", req.Page)
	return &pb.FindAllResponse{}, nil
}
