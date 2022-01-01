package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Ja7ad/grpcPractices/pb"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

const (
	defaultServerPort = 9911
)

var (
	port = flag.Int("port", defaultServerPort, "gRPC server port")
)

type server struct {
	pb.UnimplementedBalanceServiceServer
}

type accountsBalance []struct {
	Id      uint32
	Name    string
	Balance []struct {
		Type string
		Fund float64
	}
}

func main() {
	log.Println("starting balance service...")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Listen Requests from address %v", lis.Addr())
	s := grpc.NewServer()
	pb.RegisterBalanceServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}

// loadBalanceFromFile load balance accounts from file for testing in server
func loadBalanceFromFile() (accountsBalance, error) {
	var balance accountsBalance

	jsonFile, err := os.Open("/home/javad/Work/Project/go/Practices/grpcPractices/server/balance.json")
	if err != nil {
		return nil, err
	}

	byteJson, _ := ioutil.ReadAll(jsonFile)

	if err := json.Unmarshal(byteJson, &balance); err != nil {
		return nil, err
	}

	return balance, nil
}

// GetBalance is implemented for rpc function
func (s *server) GetBalance(in *pb.BalanceReq, stream pb.BalanceService_GetBalanceServer) error {
	balance, err := loadBalanceFromFile()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("New request for ID %v", in.GetId())

	for _, b := range balance {
		if b.Id == in.GetId() {
			for _, balanceItems := range b.Balance {
				res := &pb.BalanceResp{Id: b.Id, Name: b.Name, Type: balanceItems.Type, Balance: balanceItems.Fund}
				stream.Send(res)
				time.Sleep(1000 * time.Millisecond)
			}
		}
	}

	return nil
}
