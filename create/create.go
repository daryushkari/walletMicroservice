package create

import (
	"context"
	"log"
	walletPB "walletMicroservice/api/pb/proto"
)

type Server struct {
}

func (s *Server) Create(ctx context.Context, in *walletPB.CreateWalletRequest) (*walletPB.MessageResponse, error) {
	log.Printf("Receive message body from client: %s", in)
	return &walletPB.MessageResponse{}, nil
}
