package main

import (
	"context"
	"flag"
	pb "github.com/Ja7ad/greeting/protos"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	defaultName = "Javad" // if flag name is empty set default name
)

var (
	address = flag.String("address", "localhost:9911", "gRPC server address and port")
	name    = flag.String("name", defaultName, "your name for greeting")
)

// rpcConnector is object grpc client
type rpcConnector struct {
	gr *grpc.ClientConn
}

// newGRPC is a constructor for create grpc client
func (r *rpcConnector) newGRPC(addr string) (*rpcConnector, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &rpcConnector{gr: conn}, nil
}

func main() {
	flag.Parse()

	// create new object for create new rpc client
	client := &rpcConnector{}
	con, err := client.newGRPC(*address)
	if err != nil {
		log.Fatalln(err)
	}

	// close grpc client after finish
	defer con.gr.Close()

	// create new greeter object for client
	c := pb.NewGreeterClient(con.gr)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// send request to server for get greeting
	r, err := c.SayHello(ctx, &pb.GreetRequest{Name: *name})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Greeting Message is : %s", r.GetMessage())
}
