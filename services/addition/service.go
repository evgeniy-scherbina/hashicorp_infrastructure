package addition

import (
	"context"

	pb "github.com/evgeniy-scherbina/hashicorp_infrastructure/pb/addition"
)

type additionServer struct {}

func NewAdditionServer() *additionServer {
	return &additionServer{}
}

func (s *additionServer) Add(_ context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	return &pb.AddResponse{
		Rez: req.A + req.B,
	}, nil
}