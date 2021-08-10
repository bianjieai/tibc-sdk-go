package client

import (
	"context"

	"github.com/irisnet/core-sdk-go/common/codec"
	"github.com/irisnet/core-sdk-go/common/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)

type queryStateClient struct {
	sdk.BaseClient
	codec.Marshaler
}

func NewQueryStateClient(bc sdk.BaseClient, cbc codec.Marshaler) queryStateClient {
	return queryStateClient{
		BaseClient: bc,
		Marshaler:  cbc,
	}
}

func (c queryStateClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

// GetClientState queries an IBC light client.
func (c queryStateClient) GetClientState(chainName string) (*QueryClientStateResponse, error) {

	in := &QueryClientStateRequest{
		ChainName: chainName,
	}

	conn, err := c.GenConn()
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).ClientState(
		context.Background(),
		in,
	)
	if err != nil {
		return &QueryClientStateResponse{}, sdk.Wrap(err)
	}

	// todo ? change res to value?
	return res, nil

}

// GetClientStates queries all the IBC light clients of a chain.
func (c queryStateClient) GetClientStates() (*QueryClientStatesResponse, error) {
	in := &QueryClientStatesRequest{}
	conn, err := c.GenConn()
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	res, err := NewQueryClient(conn).ClientStates(
		context.Background(),
		in,
	)
	// todo ? change res to value?

	return res, err
}

// GetConsensusState queries a consensus state associated with a client state at
// a given height.
func (c queryStateClient) GetConsensusState(chainName string, height uint64) (*QueryConsensusStateResponse, error) {
	req := &QueryConsensusStateRequest{
		ChainName:      chainName,
		RevisionHeight: height,
	}
	conn, err := c.GenConn()
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).ConsensusState(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	// todo ? change res to value?

	return res, nil
}

// ConsensusStates queries all the consensus state associated with a given
// client.
func (c queryStateClient) ConsensusStates(chainName string) (*QueryConsensusStatesResponse, error) {
	req := &QueryConsensusStatesRequest{
		ChainName: chainName,
	}
	conn, err := c.GenConn()
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).ConsensusStates(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	// todo ? change res to value?

	return res, nil
}

// Relayers queries all the relayers associated with a given
// client.
func (c queryStateClient) Relayers(chainName string) (*QueryRelayersResponse, error) {
	req := &QueryRelayersRequest{
		ChainName: chainName,
	}
	conn, err := c.GenConn()
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).Relayers(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	// todo ? change res to value?

	return res, nil
}
func (c queryStateClient) UpdateClient(msgUpdateClient UpdateClientRequest) (*MsgUpdateClientResponse, error) {
	req := &MsgUpdateClient{}
	conn, err := c.GenConn()
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	res, err := NewMsgClient(conn).UpdateClient(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	// todo ? change res to value?
	return res, nil
}
