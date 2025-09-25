package main

import (
	"io"
	"log"

	pb "github.com/ayushhhhyqxv/grpc-server/proto"
)

func (s *command) InputBidirectionalStreaming(stream pb.GreetService_InputBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		
		log.Printf("Received from bi-directional client: %v", req.Name)
		
		
		res := &pb.Response{
			Message: "Voila " + req.Name,
		}
		
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}