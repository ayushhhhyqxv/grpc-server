package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ayushhhhyqxv/grpc-server/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameList){
	log.Printf("Streaming Starting ! ")
	stream,err := client.InputServerStreaming(context.Background(),names)
	if err!=nil {
		log.Fatalf("Could'nt display names: %v",err)
	}

	for {
		message,err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err!= nil {
			log.Fatalf("Error while receiving: %v",err)
		}
		log.Println(message)
	}
	log.Println("Streaming Ended !")
}