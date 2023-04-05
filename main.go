package main

import (
	"log"
	"net"
	walletPB "walletMicroservice/api/pb/proto"
	"walletMicroservice/delivery/grpc"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := delivery.Server{}

	grpcServer := grpc.NewServer()

	walletPB.RegisterWalletServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
