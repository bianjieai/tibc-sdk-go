package client

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client/flags"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	"github.com/spf13/pflag"
)

type Client struct {
	QueryClient
	MsgClient
	*query.PageRequest

	chainName string
	delay     uint64
}

type Config struct {
	chainName string
	delay     uint64
}

func NewQueryClientForClient(cc grpc1.ClientConn, config Config) (*Client, error) {
	var flagSet *pflag.FlagSet
	pageKey, _ := flagSet.GetString(flags.FlagPageKey)
	offset, _ := flagSet.GetUint64(flags.FlagOffset)
	limit, _ := flagSet.GetUint64(flags.FlagLimit)
	countTotal, _ := flagSet.GetBool(flags.FlagCountTotal)
	page, _ := flagSet.GetUint64(flags.FlagPage)

	if page > 1 && offset > 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "page and offset cannot be used together")
	}

	if page > 1 {
		offset = (page - 1) * limit
	}

	req := &query.PageRequest{
		Key:        []byte(pageKey),
		Offset:     offset,
		Limit:      limit,
		CountTotal: countTotal,
	}
	return &Client{
		QueryClient: NewQueryClient(cc),
		MsgClient:   NewMsgClient(cc),
		PageRequest: req,
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
		Pagination: c.PageRequest,
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
		Pagination: c.PageRequest,
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
