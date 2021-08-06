package client

import (
	"context"

	grpc1 "github.com/gogo/protobuf/grpc"
	"google.golang.org/grpc"
)

type Client struct {
	QueryClient

	chainName string
	delay     uint64
}

func NewQueryClientForClient(cc grpc1.ClientConn, config ...string) (*Client, error) {
	a := new(Client)
	a.QueryClient = NewQueryClient(cc)
	// chainName and delay
	return a, nil
}

// GetClientState getclientstate
func (c *Client) GetClientState(ctx context.Context, in *QueryClientStateRequest, opts ...grpc.CallOption) (*QueryClientStateResponse, error) {
	return c.QueryClient.ClientState(ctx, in, opts...)
}

// GetClientStates getclientstates
func (c *Client) GetClientStates(ctx context.Context, in *QueryClientStatesRequest, opts ...grpc.CallOption) (*QueryClientStatesResponse, error) {
	return c.QueryClient.ClientStates(ctx, in, opts...)
}

// GetConsensusState getconsensusstate
func (c *Client) GetConsensusState(ctx context.Context, in *QueryConsensusStateRequest, opts ...grpc.CallOption) (*QueryConsensusStateResponse, error) {
	return c.QueryClient.ConsensusState(ctx, in, opts...)
}

// GetConsensusStates getconsensusstates
func (c *Client) GetConsensusStates(ctx context.Context, in *QueryConsensusStatesRequest, opts ...grpc.CallOption) (*QueryConsensusStatesResponse, error) {
	return c.QueryClient.ConsensusStates(ctx, in, opts...)
}

// GetRelayers getrelayers
func (c *Client) GetRelayers(ctx context.Context, in *QueryRelayersRequest, opts ...grpc.CallOption) (*QueryRelayersResponse, error) {
	return c.QueryClient.Relayers(ctx, in, opts...)
}
