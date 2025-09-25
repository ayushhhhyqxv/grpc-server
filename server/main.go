package main

import (
	"log"
	"net"
	pb "github.com/ayushhhhyqxv/grpc-server/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type command struct {
	pb.UnimplementedGreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server : %v",err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer,&command{})
	if err := grpcServer.Serve(lis);err!=nil{
		log.Fatalf("Failed to Start : %v",err)
	}

}