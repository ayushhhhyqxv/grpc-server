package main

import (
	"context"
	"log"
	"time"
	"io"

	pb "github.com/ayushhhhyqxv/grpc-server/proto"
)

func callSayHelloBiDirectionalStream(client pb.GreetServiceClient, names *pb.NameList){
	log.Printf("Bi-Directional Client Side Streaming Started !")
	stream,err := client.InputBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could'nt Send Names: %v",err)
	}
	waitc := make(chan struct{})

	go func(){ // receives streaming data concurrently from server
		for{
			message,err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while Streaming")
			}
			log.Print("Received From Server")
			log.Println(message)
		}
		close(waitc)
	}()

	for _,name := range names.Names {
		req := &pb.Request{
			Name : name,
		}
		if err:= stream.Send(req);err!=nil{
			log.Fatalf("Error While Sending %v",err)
		}
		time.Sleep(2*time.Second)

	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bi-Directional Streaming Finished ! ")


}