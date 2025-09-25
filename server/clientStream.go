package main

import (
	"io"
	"log"

	pb "github.com/ayushhhhyqxv/grpc-server/proto"
)


func (s *command) InputClientStreaming(stream pb.GreetService_InputClientStreamingServer ) error {
	var messages []string
	for {
		req,err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Messages: messages})
		}
		if err!=nil {
			return err
		}
		log.Printf("Got the Request -> %v",req.Name)
		messages = append(messages, "Namaskar! "+req.Name)
	}
}