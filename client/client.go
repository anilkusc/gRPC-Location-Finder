package main

import (
	"context"
	"log"

	pb "github.com/anilkusc/gRPC-Location-Finder/protos"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := pb.NewLocationDeliveryClient(conn)
	response, err := c.Deliver(context.Background())
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	waitc := make(chan struct{})
	msg := &pb.Coordinates{Ip: "127.0.0.1"}
	go func() {
		for {
			log.Println("Sending msg...")
			response.Send(msg)
		}
	}()
	<-waitc
	response.CloseSend()
}
