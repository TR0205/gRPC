package main

import (
	"context"
	"log"
	"time"

	pb "example.com/grpc/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Printf("doLongGreet was invoked")

	// GreetRequestのポインタを複数持つ配列を定義
	reqs := []*pb.GreetRequest{
		{FirstName: "aaaaa"},
		{FirstName: "bbbbbb"},
		{FirstName: "cccccccccc"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatal("Error while calling LongGreet %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatal("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
