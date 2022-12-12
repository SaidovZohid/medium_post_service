package service

import (
	"context"
	"time"

	pb "github.com/SaidovZohid/medium_post_service/genproto/post_service"
	"github.com/SaidovZohid/medium_post_service/storage"
	"github.com/SaidovZohid/medium_post_service/storage/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PostService struct {
	pb.UnimplementedPostServiceServer
	storage storage.StorageI
}

func NewPostService(strg *storage.StorageI) *PostService {
	return &PostService{
		storage: *strg,
	}
}

func (s *PostService) Create(ctx context.Context, req *pb.Post) (*pb.Post, error) {
	post, err := s.storage.Post().Create(&repo.Post{
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
		UserID:      req.UserId,
		CategoryID:  req.CategoryId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &pb.Post{
		Id:          post.ID,
		Title:       post.Title,
		Description: post.Description,
		ImageUrl:    post.ImageUrl,
		UserId:      post.UserID,
		CategoryId:  post.CategoryID,
		CreatedAt:   post.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   post.UpdatedAt.Format(time.RFC3339),
		ViewsCount:  int64(post.ViewsCount),
	}, nil
}

func (s *PostService) Get(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {
	post, err := s.storage.Post().Get(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &pb.Post{
		Id:          post.ID,
		Title:       post.Title,
		Description: post.Description,
		ImageUrl:    post.ImageUrl,
		UserId:      post.UserID,
		CategoryId:  post.CategoryID,
		CreatedAt:   post.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   post.UpdatedAt.Format(time.RFC3339),
		ViewsCount:  int64(post.ViewsCount),
	}, nil
}

func (s *PostService) Update(ctx context.Context, req *pb.Post) (*pb.Post, error) {
	post, err := s.storage.Post().Update(&repo.Post{
		ID:          req.Id,
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
		UserID:      req.UserId,
		CategoryID:  req.CategoryId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &pb.Post{
		Id:          post.ID,
		Title:       post.Title,
		Description: post.Description,
		ImageUrl:    post.ImageUrl,
		UserId:      post.UserID,
		CategoryId:  post.CategoryID,
		CreatedAt:   post.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   post.UpdatedAt.Format(time.RFC3339),
		ViewsCount:  int64(post.ViewsCount),
	}, nil
}

func (s *PostService) Delete(ctx context.Context, req *pb.GetPostRequest) (*emptypb.Empty, error) {
	err := s.storage.Post().Delete(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &emptypb.Empty{}, nil
}