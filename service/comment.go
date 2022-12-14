package service

import (
	"context"
	"time"

	pb "github.com/SaidovZohid/medium_post_service/genproto/post_service"
	"github.com/SaidovZohid/medium_post_service/genproto/user_service"
	grpcPkg "github.com/SaidovZohid/medium_post_service/pkg/grpc_client"
	"github.com/SaidovZohid/medium_post_service/storage"
	"github.com/SaidovZohid/medium_post_service/storage/repo"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CommentService struct {
	pb.UnimplementedCommentServiceServer
	storage    storage.StorageI
	grpcClient grpcPkg.GrpcClientI
	logger     *logrus.Logger
}

func NewCommentService(strg *storage.StorageI, grpc grpcPkg.GrpcClientI, log *logrus.Logger) *CommentService {
	return &CommentService{
		storage:    *strg,
		grpcClient: grpc,
		logger:     log,
	}
}

func (s *CommentService) Create(ctx context.Context, req *pb.Comment) (*pb.Comment, error) {
	comment, err := s.storage.Comment().Create(&repo.Comment{
		PostID:      req.PostId,
		UserID:      req.UserId,
		Description: req.Description,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error in create comment service: %v", err)
	}

	user, err := s.grpcClient.UserService().Get(context.Background(), &user_service.IdRequest{
		Id: req.UserId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error in while getting user info comment service: %v", err)
	}

	return &pb.Comment{
		Id:          comment.ID,
		PostId:      comment.PostID,
		UserId:      comment.UserID,
		Description: comment.Description,
		CreatedAt:   comment.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   comment.UpdatedAt.Format(time.RFC3339),
		User: &pb.CommentUser{
			Id:              user.Id,
			FirstName:       user.FirstName,
			LastName:        user.LastName,
			Email:           user.Email,
			ProfileImageUrl: user.ProfileImageUrl,
		},
	}, nil
}
func (s *CommentService) Delete(ctx context.Context, req *pb.DeleteCommentRequest) (*emptypb.Empty, error) {
	err := s.storage.Comment().Delete(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internale server while deleting comment: %v", err)
	}

	return &emptypb.Empty{}, nil
}
func (s *CommentService) GetAll(ctx context.Context, req *pb.GetAllCommentsParamsReq) (*pb.GetAllCommentsResponse, error) {
	return nil, nil
}

// func (s *CommentService) Update(ctx context.Context, req *pb.Comment) (*pb.Comment, error) {
// 	comment, err := s.storage.Comment().Update(&repo.Comment{
// 		Description: req.Description,
// 		ID:          req.Id,
// 	})
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "internal server error while updating comment: %v", err)
// 	}
// 	user, err := s.grpcClient.UserService().Get(context.Background(), &user_service.IdRequest{Id: req.UserId})
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "internal server error while updating comment: %v", err)
// 	}
// 	return nil, nil
// }
