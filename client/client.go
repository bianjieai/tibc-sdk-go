package client

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	"google.golang.org/grpc"
)

type Client struct {
	QueryClient
	MsgClient
	pageReq   *query.PageRequest
	chainName string
	delay     uint64
}

type Config struct {
	chainName string
	delay     uint64
}

func NewQueryClientForClient(cc grpc1.ClientConn, config Config) (*Client, error) {
	return &Client{
		QueryClient: NewQueryClient(cc),
		MsgClient:   NewMsgClient(cc),
		chainName:   config.chainName,
		delay:       config.delay,
	}, nil
}

// GetClientState queries an IBC light client.
func (c *Client) GetClientState(ctx context.Context) (*QueryClientStateResponse, error) {
	in := &QueryClientStateRequest{
		ChainName: c.chainName,
	}
	return c.QueryClient.ClientState(ctx, in)
}

// GetClientStates queries all the IBC light clients of a chain.
func (c *Client) GetClientStates(ctx context.Context) (*QueryClientStatesResponse, error) {
	req := &QueryClientStatesRequest{
		Pagination: c.pageReq,
	}
	return c.QueryClient.ClientStates(ctx, req)
}

// GetConsensusState queries a consensus state associated with a client state at
// a given height.
func (c *Client) GetConsensusState(ctx context.Context, pageReq *query.PageRequest) (*QueryConsensusStateResponse, error) {
	req := &QueryConsensusStateRequest{
		ChainName: c.chainName,
	}
	return c.QueryClient.ConsensusState(ctx, req)
}

// ConsensusStates queries all the consensus state associated with a given
// client.
func (c *Client) ConsensusStates(ctx context.Context) (*QueryConsensusStatesResponse, error) {
	req := &QueryConsensusStatesRequest{
		ChainName:  c.chainName,
		Pagination: c.pageReq,
	}
	return c.QueryClient.ConsensusStates(ctx, req)
}

// Relayers queries all the relayers associated with a given
// client.
func (c *Client) Relayers(ctx context.Context) (*QueryRelayersResponse, error) {
	req := &QueryRelayersRequest{
		ChainName: c.chainName,
	}
	return c.QueryClient.Relayers(ctx, req)
}

func (c *Client) UpdateClient(ctx context.Context, in *MsgUpdateClient, opts ...grpc.CallOption) (*MsgUpdateClientResponse, error) {
	content, err := NewCreateClientProposal(title, description, chainName, clientState, consensusState)
	if err != nil {
		return err
	}
	return c.MsgClient.UpdateClient(ctx, in, opts...)
}
