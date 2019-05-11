package multiplication

import (
	"context"

	pb "github.com/evgeniy-scherbina/hashicorp_infrastructure/pb/multiplication"
)

type multiplicationServer struct {}

func NewMultiplicationServer() *multiplicationServer {
	return &multiplicationServer{}
}

func (s *multiplicationServer) Mul(_ context.Context, req *pb.MulRequest) (*pb.MulResponse, error) {
	return &pb.MulResponse{
		Rez: req.A * req.B,
	}, nil
}