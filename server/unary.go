package main

import ("context"
	pb "github.com/ayushhhhyqxv/grpc-server/proto")

func (s *command) Input(ctx context.Context, req *pb.NoParam) (*pb.Response, error) {
	return &pb.Response{
		Message: "Unary Response Sent", // This is the actual response data
	}, nil // nil means no error occurred
}	