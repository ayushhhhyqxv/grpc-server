package main

import (
	"log"
	"time"

	pb "github.com/ayushhhhyqxv/grpc-server/proto"
)

func (s *command) InputServerStreaming(req *pb.NameList, stream pb.GreetService_InputServerStreamingServer) error {
	log.Printf("Names Passed: %v",req.Names)
	for _,name := range req.Names {
		res := &pb.Response{
			Message: "Hello "+ name,
		}
		if err:= stream.Send(res); err!=nil{
			return err
		}
		time.Sleep(2*time.Second) // To see Simulation ! 
	}
	return nil
}