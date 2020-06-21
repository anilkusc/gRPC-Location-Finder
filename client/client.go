package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"flag"

	pb "github.com/anilkusc/gRPC-Location-Finder/protos"
	"google.golang.org/grpc"
)

type Client struct {
	ip string
	x  int32
	y  int32
}

var (
	clients []Client
	ip      = flag.String("ip", "192.168.1.1", "Client IP")
	x       = flag.Int("x", 10, "Client X Coordinate")
	y       = flag.Int("y", 10, "Client Y Coordinate")
)

func main() {
	flag.Parse()
	var me, others Client

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := pb.NewLocationDeliveryClient(conn)

	me.ip = *ip
	me.x = int32(*x)
	me.y = int32(*y)
	req := pb.Client{Ip: me.ip, X: me.x, Y: me.y}
	stream, _ := c.Deliver(context.Background(), &req)
	log.Printf("%s sent", req.Ip)
	clients = nil
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		others.ip = resp.Ip
		others.x = resp.X
		others.y = resp.Y
		clients = append(clients, others)
	}

	log.Printf("finished")
	fmt.Println(clients)
}
