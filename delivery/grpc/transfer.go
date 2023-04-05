package delivery

import (
	"context"
	"log"
	walletPB "walletMicroservice/api/pb/proto"
)

type Server struct {
	walletPB.UnimplementedWalletServer
}

func (s *Server) Create(ctx context.Context, in *walletPB.CreateWalletRequest) (*walletPB.MessageResponse, error) {
	log.Printf("Receive message body from client: %s", in)
	return &walletPB.MessageResponse{}, nil
}

func (s *Server) Charge(ctx context.Context, in *walletPB.ChangeWalletBalance) (*walletPB.MessageResponse, error) {
	log.Printf("Receive message body from client: %s", in)
	return &walletPB.MessageResponse{}, nil
}

func (s *Server) Refund(ctx context.Context, in *walletPB.ChangeWalletBalance) (*walletPB.MessageResponse, error) {
	log.Printf("Receive message body from client: %s", in)
	return &walletPB.MessageResponse{}, nil
}

func (s *Server) InitialTransfer(ctx context.Context, in *walletPB.TransferRequest) (*walletPB.InitialTransferResponse, error) {
	log.Printf("Receive message body from client: %s", in)
	return &walletPB.InitialTransferResponse{}, nil
}

func (s *Server) ConfirmTransfer(ctx context.Context, in *walletPB.DoTransfer) (*walletPB.MessageResponse, error) {
	log.Printf("Receive message body from client: %s", in)
	return &walletPB.MessageResponse{}, nil
}

func (s *Server) CancelTransfer(ctx context.Context, in *walletPB.DoTransfer) (*walletPB.MessageResponse, error) {
	log.Printf("Receive message body from client: %s", in)
	return &walletPB.MessageResponse{}, nil
}
