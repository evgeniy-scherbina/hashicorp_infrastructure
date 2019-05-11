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

	pb "github.com/evgeniy-scherbina/hashicorp_infrastructure/pb/addition"
	"github.com/evgeniy-scherbina/hashicorp_infrastructure/services/addition"
)

const (
	defaultRpcListen = "0.0.0.0:9090"
)

func main() {
	log.Info("Addition Service Starting")

	var g run.Group

	// GRPC endpoints
	{
		grpcServer := grpc.NewServer()

		r := load.NewRateReporter(time.Minute)
		lbpb.RegisterLoadReportServer(grpcServer, r)

		additionServer := addition.NewAdditionServer()
		pb.RegisterAdditionServiceServer(grpcServer, additionServer)

		g.Add(func() error {
			log.Info("Start Addition GRPC endpoints")

			lis, err := net.Listen("tcp", defaultRpcListen)
			if err != nil {
				return fmt.Errorf("failed to listen: %v", err)
			}

			return grpcServer.Serve(lis)
		}, func(err error) {
			log.Info("Stop Addition GRPC endpoints")
			grpcServer.GracefulStop()
		})
	}

	log.Infof("The addition-service was terminated with: %v", g.Run())
}
