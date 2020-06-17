package main

import (
	"io"
	"log"
	"net"

	pb "github.com/anilkusc/gRPC-Location-Finder"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type server struct{}

func (s server) Login(ctx context.Context, request *sso.LoginRequest) (string) {
	p, _ := peer.FromContext(ctx)
	request.Frontendip = p.Addr.String()
	}

func (s server) Deliver(stream pb.LocationDelivery_DeliverServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		key := serialize(in.Location)
                ... // look for notes to be sent to client
		for _, note := range s.routeNotes[key] {
			if err := stream.Send(note); err != nil {
				return err
			}
		}
	}
}

func main() {
	// create listiner
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()
	pb.RegisterMathServer(s, server{})

	// and start...
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
