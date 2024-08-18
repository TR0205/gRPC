package main

import (
	"fmt"
	"log"

	pb "example.com/grpc/greet/proto"
	"google.golang.org/grpc"
)

// inはGreetRequestのポインタを受け取っている
// ->構造体を値として渡すよりも効率が良いため
func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream grpc.ServerStreamingServer[pb.GreetRespponse]) error {
	log.Printf("Greet many times function was invoked with %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)

		// GreetRespponsem構造体を新しく生成し、ポインタを渡す
		stream.Send(&pb.GreetRespponse{
			Result: res,
		})
	}

	return nil
}
