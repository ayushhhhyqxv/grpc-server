package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ayushhhhyqxv/grpc-server/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList){
	log.Printf("Client-Side Streaming Started !")
	stream, err := client.InputClientStreaming(context.Background())
	if err!=nil {
		log.Fatalf("Could not send Names : %v",err)
	}

	for _,name := range names.Names{
		req := &pb.Request{
			Name : name,
		}
		if err := stream.Send(req);err!=nil {
			log.Fatalf("Error while Sending from Client Side: %v",err)
		}
		log.Printf("Sent the request with Name: %s",name)
		time.Sleep(2*time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client Streaming Finished !")
	if err!=nil {
		log.Fatalf("Error while receiving in server-Side : %v",err)
	}
	log.Printf("%v",res.Messages)
}


