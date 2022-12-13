package service

import (
	"context"

	pb "github.com/SaidovZohid/medium_post_service/genproto/post_service"
	"github.com/SaidovZohid/medium_post_service/storage"
	"github.com/SaidovZohid/medium_post_service/storage/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LikeService struct {
	pb.UnimplementedLikeServiceServer
	storage storage.StorageI
}

func NewLikeService(strg *storage.StorageI) *LikeService {
	return &LikeService{
		storage: *strg,
	}
}

func (s *LikeService) CreateOrUpdate(ctx context.Context, req *pb.Like) (*pb.Like, error) {
	like, err := s.storage.Like().CreateOrUpdate(&repo.Like{
		PostID: req.PostId,
		UserID: req.UserId,
		Status: req.Status,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return parseLike(like), nil
}

func (s *LikeService) Get(ctx context.Context, req *pb.GetLikeRequest) (*pb.Like, error) {
	like, err := s.storage.Like().Get(req.UserId, req.PostId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return parseLike(like), nil
}

func (s *LikeService) GetLikesDislikesCount(ctx context.Context, req *pb.GetLikesRequest) (*pb.LikesDislikesCountResponse, error) {
	counts, err := s.storage.Like().GetLikesDislikesCount(req.PostId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &pb.LikesDislikesCountResponse{
		Likes:    counts.Likes,
		Dislikes: counts.Dislikes,
	}, nil
}

func parseLike(req *repo.Like) *pb.Like {
	return &pb.Like{
		Id:     req.ID,
		UserId: req.UserID,
		PostId: req.PostID,
		Status: req.Status,
	}
}
