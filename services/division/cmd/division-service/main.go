package main

import (
	"fmt"
	"net"
	"time"

	lbpb "github.com/bsm/grpclb/grpclb_backend_v1"
	"github.com/bsm/grpclb/load"
	"github.com/oklog/run"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "github.com/evgeniy-scherbina/hashicorp_infrastructure/pb/division"
	"github.com/evgeniy-scherbina/hashicorp_infrastructure/services/division"
)

const (
	defaultRpcListen = "0.0.0.0:9090"
)

func main() {
	log.Info("Division Service Starting")

	var g run.Group

	// GRPC endpoints
	{
		grpcServer := grpc.NewServer()

		r := load.NewRateReporter(time.Minute)
		lbpb.RegisterLoadReportServer(grpcServer, r)

		divisionServer := addition.NewDivisionServer()
		pb.RegisterDivisionServiceServer(grpcServer, divisionServer)

		g.Add(func() error {
			log.Info("Start Division GRPC endpoints")

			lis, err := net.Listen("tcp", defaultRpcListen)
			if err != nil {
				return fmt.Errorf("failed to listen: %v", err)
			}

			return grpcServer.Serve(lis)
		}, func(err error) {
			log.Info("Stop Division GRPC endpoints")
			grpcServer.GracefulStop()
		})
	}

	log.Infof("The division-service was terminated with: %v", g.Run())
}
