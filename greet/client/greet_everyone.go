package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "example.com/grpc/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Printf("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "aaaaa"},
		{FirstName: "bbbbbb"},
		{FirstName: "cccccccccc"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			// リクエスト終了
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal("Error while reading client stream: %v\n", err)
			}

			log.Printf("Received: %v\n", res.Result)
		}
		close(waitc)
	}()
	<-waitc
}
