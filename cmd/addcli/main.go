package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	addPb "github.com/evgeniy-scherbina/hashicorp_infrastructure/pb/addition"
	divPb "github.com/evgeniy-scherbina/hashicorp_infrastructure/pb/division"
	subPb "github.com/evgeniy-scherbina/hashicorp_infrastructure/pb/subtraction"
	mulPb "github.com/evgeniy-scherbina/hashicorp_infrastructure/pb/multiplication"
	addClient "github.com/evgeniy-scherbina/hashicorp_infrastructure/services/addition/client"
	subClient "github.com/evgeniy-scherbina/hashicorp_infrastructure/services/subtraction/client"
	mulClient "github.com/evgeniy-scherbina/hashicorp_infrastructure/services/multiplication/client"
	divClient "github.com/evgeniy-scherbina/hashicorp_infrastructure/services/division/client"
)

const grpclbAddress = "grpclb:8383"

func main() {
	//conn, err := grpc.Dial("addition:9090", grpc.WithInsecure())
	//if err != nil {
	//	log.Fatal(err)
	//}

	// grpcClient := pb.NewAdditionServiceClient(conn)
	ctxb := context.Background()

	// ----------------------------------------------------------------------------------------------------
	{
		grpcClient, err := addClient.NewAdditionServiceClient(ctxb, grpclbAddress)
		if err != nil {
			log.Fatal(err)
		}

		req := addPb.AddRequest{
			A: 2,
			B: 3,
		}
		resp, err := grpcClient.Add(ctxb, &req, grpc.FailFast(false))
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp.Rez)
	}

	// ----------------------------------------------------------------------------------------------------
	{
		grpcClient, err := subClient.NewSubtractionServiceClient(ctxb, grpclbAddress)
		if err != nil {
			log.Fatal(err)
		}

		req := subPb.SubRequest{
			A: 2,
			B: 3,
		}
		resp, err := grpcClient.Sub(ctxb, &req, grpc.FailFast(false))
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp.Rez)
	}

	// ----------------------------------------------------------------------------------------------------
	{
		grpcClient, err := mulClient.NewMultiplicationServiceClient(ctxb, grpclbAddress)
		if err != nil {
			log.Fatal(err)
		}

		req := mulPb.MulRequest{
			A: 2,
			B: 3,
		}
		resp, err := grpcClient.Mul(ctxb, &req, grpc.FailFast(false))
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp.Rez)
	}

	// ----------------------------------------------------------------------------------------------------
	{
		grpcClient, err := divClient.NewDivisionServiceClient(ctxb, grpclbAddress)
		if err != nil {
			log.Fatal(err)
		}

		req := divPb.DivRequest{
			A: 2,
			B: 3,
		}
		resp, err := grpcClient.Div(ctxb, &req, grpc.FailFast(false))
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp.Rez)
	}

	// ----------------------------------------------------------------------------------------------------
}
