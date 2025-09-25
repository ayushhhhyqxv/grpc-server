package main

import (
	"log"
	pb "github.com/ayushhhhyqxv/grpc-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost" + port,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil {
		log.Fatalf("Did'nt Connect : %v",err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	// callSayHello(client) // unary client 

	names := &pb.NameList{
		Names: []string{
		"Ayush",
		"Yash",
		"Kartik",
	},
	}

	// callSayHelloServerStream(client,names)
	// callSayHelloClientStream(client,names)
	callSayHelloBiDirectionalStream(client,names)


	


}