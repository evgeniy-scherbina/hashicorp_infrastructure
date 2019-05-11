package client

import (
	"context"
	"fmt"
	"time"

	"github.com/bsm/grpclb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "github.com/evgeniy-scherbina/hashicorp_infrastructure/pb/subtraction"
)

const ServiceName = "subtraction-service"

// NewSubtractionServiceClient returns a grpc client to the subtraction service
func NewSubtractionServiceClient(ctx context.Context, grpclbAddr string) (pb.SubServiceClient, error) {
	resolver := grpclb.NewResolver(&grpclb.Options{
		Address: grpclbAddr,
	})

	connCtx, _ := context.WithTimeout(ctx, 10*time.Second)
	conn, err := grpc.DialContext(
		connCtx,
		fmt.Sprintf("%s-rpc", ServiceName),
		grpc.WithInsecure(),
		grpc.WithBalancer(grpc.RoundRobin(resolver)),
	)
	if err != nil {
		return nil, err
	}

	go func() {
		<-ctx.Done()
		if err := conn.Close(); err != nil {
			grpclog.Infof("Failed to close connection: %v", err)
		}
	}()

	return pb.NewSubServiceClient(conn), nil
}
