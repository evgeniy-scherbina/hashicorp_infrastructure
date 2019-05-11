package subtraction

import (
	"context"

	pb "github.com/evgeniy-scherbina/hashicorp_infrastructure/pb/subtraction"
)

type subtractionServer struct {}

func NewSubtractionServer() *subtractionServer {
	return &subtractionServer{}
}

func (s *subtractionServer) Sub(_ context.Context, req *pb.SubRequest) (*pb.SubResponse, error) {
	return &pb.SubResponse{
		Rez: req.A - req.B,
	}, nil
}