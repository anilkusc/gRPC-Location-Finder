package main

import (
	"log"
	"net"

	pb "github.com/anilkusc/gRPC-Location-Finder/protos"

	"google.golang.org/grpc"
)

type Client struct {
	ip string
	x  int32
	y  int32
}

var myclients []Client

type server struct{}

func (s server) Deliver(client *pb.Client, stream pb.LocationDelivery_DeliverServer) error {

	//for {

	var myclient Client
	myclient.ip = client.GetIp()
	myclient.x = client.GetX()
	myclient.y = client.GetY()
	log.Printf("Client info will be added ip,x,y: %s , %d , %d \n", myclient.ip, myclient.x, myclient.y)
	if myclients == nil {
		myclients = append(myclients, myclient)
	} else {
		Sync(myclient)
	}
	for _, oneclient := range myclients {

		if err := stream.Send(&pb.Client{Ip: oneclient.ip, X: oneclient.x, Y: oneclient.y}); err != nil {
			log.Printf("send error %v", err)
			continue
		}
	}
	log.Printf("sended")
	//}
	return nil
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
