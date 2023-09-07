package client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Jimbo8702/toll-calc/types"
)

type GRPCClient struct {
	Endpoint string
	client 	 types.AggregatorClient
}

func NewGRPCCLient(endpoint string) (*GRPCClient, error) {
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := types.NewAggregatorClient(conn)
	return &GRPCClient{
		Endpoint: endpoint,
		client: c,
	}, nil
}
//TODO: implement GetInvoice
func (c *GRPCClient) Aggregate(ctx context.Context, req *types.AggregateRequest) error {
	_, err := c.client.Aggregate(ctx, req)
	return err
}