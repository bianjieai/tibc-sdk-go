package tendermint

import (
	"context"

	commoncodec "github.com/irisnet/core-sdk-go/common/codec"
	cryptotypes "github.com/irisnet/core-sdk-go/common/codec/types"
	"github.com/irisnet/core-sdk-go/types"
)

type tendermintClient struct {
	types.BaseClient
	commoncodec.Marshaler
}


func NewClient(bc types.BaseClient,cdc commoncodec.Marshaler) *tendermintClient {
	return &tendermintClient{
		BaseClient: bc,
		Marshaler:  cdc,
	}
}


func (c tendermintClient) RegisterInterfaceTypes(registry cryptotypes.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

// GetClientState queries an IBC light client.
func (c tendermintClient) GetClientState(chainName string) (*QueryClientStateResponse, error) {

	in := &QueryClientStateRequest{
		ChainName: chainName,
	}

	conn, err := c.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}

	res, err := NewQueryClient(conn).ClientState(
		context.Background(),
		in,
	)
	if err != nil {
		return &QueryClientStateResponse{}, types.Wrap(err)
	}
	var clientState ClientState
	if err := c.Marshaler.UnpackAny(res.ClientState, &clientState); err != nil {
		return nil, types.Wrap(err)
	}
	// todo ? change res to value?
	return res, nil

}

// GetClientStates queries all the IBC light clients of a chain.
func (c tendermintClient) GetClientStates() (*QueryClientStatesResponse, error) {
	in := &QueryClientStatesRequest{}
	conn, err := c.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
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
func (c tendermintClient) GetConsensusState(chainName string, height uint64) (*QueryConsensusStateResponse, error) {
	req := &QueryConsensusStateRequest{
		ChainName:      chainName,
		RevisionHeight: height,
	}
	conn, err := c.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}

	res, err := NewQueryClient(conn).ConsensusState(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, types.Wrap(err)
	}
	// todo ? change res to value?

	return res, nil
}

// ConsensusStates queries all the consensus state associated with a given
// client.
func (c tendermintClient) ConsensusStates(chainName string) (*QueryConsensusStatesResponse, error) {
	req := &QueryConsensusStatesRequest{
		ChainName: chainName,
	}
	conn, err := c.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}

	res, err := NewQueryClient(conn).ConsensusStates(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, types.Wrap(err)
	}
	// todo ? change res to value?

	return res, nil
}

// Relayers queries all the relayers associated with a given
// client.
func (c tendermintClient) Relayers(chainName string) (*QueryRelayersResponse, error) {
	req := &QueryRelayersRequest{
		ChainName: chainName,
	}
	conn, err := c.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}

	res, err := NewQueryClient(conn).Relayers(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, types.Wrap(err)
	}
	// todo ? change res to value?

	return res, nil
}
func (c tendermintClient) UpdateClient(msgUpdateClient UpdateClientRequest) (*MsgUpdateClientResponse, error) {
	req := &MsgUpdateClient{}
	conn, err := c.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}
	res, err := NewMsgClient(conn).UpdateClient(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, types.Wrap(err)
	}
	// todo ? change res to value?
	return res, nil
}
