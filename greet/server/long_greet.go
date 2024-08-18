package main

import (
	"fmt"
	"io"
	"log"

	pb "example.com/grpc/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) LongGreet(stream grpc.ClientStreamingServer[pb.GreetRequest, pb.GreetRespponse]) error {
	log.Printf("LongGreet function was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		// リクエスト終了
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetRespponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatal("Error while reading client stream: %v\n", err)
		}

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}
