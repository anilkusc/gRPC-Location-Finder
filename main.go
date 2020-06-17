package main

import (
	"io"
	"log"
	"net"

	pb "github.com/anilkusc/gRPC-Location-Finder/protos"

	"google.golang.org/grpc"
)

type server struct{}

func (s server) Deliver(stream pb.LocationDelivery_DeliverServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		ip := in.Ip
		log.Println("Client ip is:", ip)
		if err := stream.Send(&pb.Coordinates{Ip: ip}); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterLocationDeliveryServer(s, server{})
	log.Println("Listening on port: 5000")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
