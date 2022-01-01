package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/Ja7ad/greeting/protos"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	defaultServerPort = 9911
)

var (
	port = flag.Int("port", defaultServerPort, "gRPC server port")
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.GreetResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", *port))
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("server listen at address %v", lis.Addr())

	if errRPC := s.Serve(lis); errRPC != nil {
		log.Fatalln(errRPC)
	}
}
