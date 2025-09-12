package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ayushhhhyqxv/grpc-server/proto"
)

func callSayHello(client pb.GreetServiceClient){
	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	res,err := client.Input(ctx,&pb.NoParam{})
	if err!=nil {
		log.Fatalf("Could not Greet You :( %v",err)
	}
	log.Printf("Data Passed: %s",res.Message)
}