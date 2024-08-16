package main

import (
	"context"
	"log"

	pb "example.com/grpc/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetRespponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.GreetRespponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
