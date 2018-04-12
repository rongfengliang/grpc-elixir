package main

import (
	"context"
	"log"

	pb "github.com/rongfengliang/grpc-elixir/proto"
	grpc "google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("connect error")
	}
	client := pb.NewGreeterClient(con)
	result, err := client.SayHello(context.Background(), &pb.HelloRequest{
		Name: "dalongdemo",
	})
	if err != nil {
		log.Println("request is error")
	}
	log.Println(result.Message)

}
