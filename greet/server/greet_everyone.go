package main

import (
	"io"
	"log"

	pb "example.com/grpc/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) GreetEveryone(stream grpc.BidiStreamingServer[pb.GreetRequest, pb.GreetRespponse]) error {
	log.Printf("Greet everyone function was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatal("Error while reading client stream: %v\n", err)
		}

		// リクエストパラメータのFirstNameを加工しレスポンス
		res := "Hello " + req.FirstName + "!"
		err = stream.Send(&pb.GreetRespponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}
