package handler

import (
	"context"
	"log"
	pb "micr-go/services/users/pb"
)

type UsersHandler struct{}

func (u *UsersHandler) FindAll(ctx context.Context, req *pb.FindAllRequest) (*pb.FindAllResponse, error) {
	log.Printf("Receive message %s", req.Page)
	users := []*pb.User{
		{
			Id:       1,
			Username: "Son",
			Password: "123",
			Email:    "nhuson994@gmail.com",
			Address:  "Hanoi",
			Phone:    "09613435079",
		},
		{
			Id:       1,
			Username: "Steven",
			Password: "123",
			Email:    "steven94@gmail.com",
			Address:  "HaiDuong",
			Phone:    "09613435079",
		},
	}
	data := pb.FindAllResponse{
		Status:  true,
		Message: "Get Users",
		Data:    users,
	}
	return &data, nil
}
