package addition

import (
	"context"

	pb "github.com/evgeniy-scherbina/hashicorp_infrastructure/pb/division"
)

type divisionServer struct {}

func NewDivisionServer() *divisionServer {
	return &divisionServer{}
}

func (s *divisionServer) Div(_ context.Context, req *pb.DivRequest) (*pb.DivResponse, error) {
	return &pb.DivResponse{
		Rez: req.A / req.B,
	}, nil
}