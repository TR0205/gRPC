package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "example.com/grpc/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	// gRPCサーバーへ接続するクライアントの生成(SSL設定は一旦無視)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	// protoから生成されたクライアント
	// GreetServiceで定義されたメソッドを呼び出すために使用
	c := pb.NewGreetServiceClient(conn)

	// doGreat(c)
	doGreetManyTimes(c)
}
