package service

import (
	"context"

	pb "github.com/SaidovZohid/medium_post_service/genproto/post_service"
	grpcPkg "github.com/SaidovZohid/medium_post_service/pkg/grpc_client"
	"github.com/SaidovZohid/medium_post_service/storage"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CommentService struct {
	pb.UnimplementedCommentServiceServer
	storage    *storage.StorageI
	grpcClient grpcPkg.GrpcClientI
}

func NewCommentService(strg *storage.StorageI, grpc grpcPkg.GrpcClientI) *CommentService {
	return &CommentService{
		storage:    strg,
		grpcClient: grpc,
	}
}

func (s *CommentService) Create(ctx context.Context, req *pb.Comment) (*pb.Comment, error) {
	return nil, nil
}
func (s *CommentService) Delete(ctx context.Context, req *pb.GetCommentRequest) (*emptypb.Empty, error) {
	return nil, nil
}
func (s *CommentService) GetAll(ctx context.Context, req *pb.GetAllCommentsParamsReq) (*pb.GetAllCommentsResponse, error) {
	return nil, nil
}
func (s *CommentService) Update(ctx context.Context, req *pb.Comment) (*pb.Comment, error) {
	return nil, nil
}
