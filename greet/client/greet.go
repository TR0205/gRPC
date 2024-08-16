package main

import (
	"context"
	"log"

	// /protoにある「package proto」と書かれている.goファイルの関数などを使用可能
	pb "example.com/grpc/greet/proto"
)

func doGreat(c pb.GreetServiceClient) {
	log.Println("doGreat was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Takashi",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
