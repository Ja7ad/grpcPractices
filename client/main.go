package main

import (
	"context"
	"flag"
	"github.com/Ja7ad/grpcPractices/pb"
	"google.golang.org/grpc"
	"io"
	"log"
)

const (
	defaultID = 1
)

var (
	address = flag.String("address", "localhost:9911", "gRPC server address and port")
	id      = flag.Int("id", defaultID, "your account id for get balance")
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
	log.Println("Starting client...")

	// create new object for create new rpc client
	client := &rpcConnector{}
	con, err := client.newGRPC(*address)
	if err != nil {
		log.Fatalln(err)
	}

	// close grpc client after finish
	defer con.gr.Close()

	c := pb.NewBalanceServiceClient(con.gr)
	fetchBalance(c)
}

func fetchBalance(c pb.BalanceServiceClient) {
	req := &pb.BalanceReq{Id: uint32(*id)}

	stream, err := c.GetBalance(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("Balance Acccount id %v with name %v is %v (%v)", res.Id, res.Name, res.Balance, res.Type)
	}
}
