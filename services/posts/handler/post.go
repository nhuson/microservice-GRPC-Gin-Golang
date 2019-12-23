package handler

import (
	"context"
	"errors"
	pb "micr-go/services/posts/pb"
	"micr-go/services/posts/repo"
)

type PostHandler struct{}

var postRepo repo.Post

func (p *PostHandler) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	postReq := req.GetPost()
	_, err := postRepo.CreatePost(ctx, repo.PostItem{
		Title:       postReq.GetTitle(),
		Description: postReq.GetDescription(),
		Slug:        postReq.GetSlug(),
		UserID:      postReq.GetUserId(),
	})
	if err != nil {
		return nil, errors.New("Create post failed")
	}

	return &pb.CreatePostResponse{
		Status:  true,
		Message: "Create post successfull!",
	}, nil
}

func (p *PostHandler) ListPost(ctx context.Context, req *pb.ListPostRequest) (*pb.ListPostResponse, error) {
	postItem := repo.PostItem{}
	cursor, _ := postRepo.ListPost(ctx)
	defer cursor.Close(ctx)
	posts := []*pb.PostResponse{}
	for cursor.Next(ctx) {
		if err := cursor.Decode(&postItem); err != nil {
			return nil, errors.New("Error decode posts")
		}
		posts = append(posts, &pb.PostResponse{
			Post: &pb.Post{
				Id:          postItem.ID.Hex(),
				Title:       postItem.Title,
				Description: postItem.Description,
				Slug:        postItem.Slug,
			},
		})
	}

	return &pb.ListPostResponse{
		Status: true,
		Data:   posts,
	}, nil
}
