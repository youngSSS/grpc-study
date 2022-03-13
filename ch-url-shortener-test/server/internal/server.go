package internal

import (
	"context"
	pb "github.com/channel-io/grpc-study/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Server struct {
	pb.RegisterAPIServer
}

func (s *Server) Ping(ctx context.Context, req *pb.Empty) (*wrapperspb.StringValue, error) {
	return wrapperspb.String("pong"), status.New(codes.OK, "").Err()
}

func (s *Server) GetRedirection(ctx context.Context, req *pb.GetRequest) (*pb.RedirectionURL, error) {
	return &pb.RedirectionURL{
		Token:       "1",
		OriginalUrl: "2",
		Description: "3",
		CreatedAt:   "4",
		ExpireAt:    "5",
		RemoveAt:    "6",
		Version:     0,
	}, status.New(codes.OK, "").Err()
}

func (s *Server) CreateRedirection(ctx context.Context, req *pb.CreateRequest) (*pb.RedirectionURL, error) {
	return &pb.RedirectionURL{
		Token:       "1",
		OriginalUrl: "2",
		Description: "3",
		CreatedAt:   "4",
		ExpireAt:    "5",
		RemoveAt:    "6",
		Version:     0,
	}, status.New(codes.OK, "").Err()
}
