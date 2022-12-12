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

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	storage storage.StorageI
}

func NewCategoryService(strg *storage.StorageI) *CategoryService {
	return &CategoryService{
		storage: *strg,
	}
}

func (s *CategoryService) Create(ctx context.Context, req *pb.Category) (*pb.Category, error) {
	category, err := s.storage.Category().Create(&repo.Category{
		Title: req.Title,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &pb.Category{
		Id:        category.ID,
		Title:     category.Title,
		CreatedAt: category.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *CategoryService) Get(ctx context.Context, req *pb.GetCategoryRequest) (*pb.Category, error) {
	category, err := s.storage.Category().Get(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &pb.Category{
		Id:        category.ID,
		Title:     category.Title,
		CreatedAt: category.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *CategoryService) Update(ctx context.Context, req *pb.Category) (*pb.Category, error) {
	category, err := s.storage.Category().Update(&repo.Category{
		ID:    req.Id,
		Title: req.Title,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &pb.Category{
		Id:        category.ID,
		Title:     category.Title,
		CreatedAt: category.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *CategoryService) Delete(ctx context.Context, req *pb.GetCategoryRequest) (*emptypb.Empty, error) {
	err := s.storage.Category().Delete(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &emptypb.Empty{}, nil
}