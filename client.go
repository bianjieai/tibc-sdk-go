package tibc_sdk_go

import (
	"context"
	"github.com/bianjieai/tibc-sdk-go/tendermint"

	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
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
	tendermint.RegisterInterfaces(registry)
}

// GetClientState queries an IBC light client.
func (c tendermintClient) GetClientState(chainName string) (tendermint.ClientState, error) {

	in := &tendermint.QueryClientStateRequest{
		ChainName: chainName,
	}

	conn, err := c.GenConn()
	if err != nil {
		return tendermint.ClientState{}, types.Wrap(err)
	}

	res, err := tendermint.NewQueryClient(conn).ClientState(
		context.Background(),
		in,
	)
	if err != nil {
		return tendermint.ClientState{}, types.Wrap(err)
	}
	var clientState tibctypes.ClientState
	if err := c.Marshaler.UnpackAny(res.ClientState, &clientState); err != nil {
		return tendermint.ClientState{}, types.Wrap(err)
	}
	// todo ? change res to value?
	return tendermint.ClientState{}, nil

}

// GetClientStates queries all the IBC light clients of a chain.
func (c tendermintClient) GetClientStates() (*tendermint.QueryClientStatesResponse, error) {
	in := &tendermint.QueryClientStatesRequest{}
	conn, err := c.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}
	res, err := tendermint.NewQueryClient(conn).ClientStates(
		context.Background(),
		in,
	)
	// todo ? change res to value?

	return res, err
}

// GetConsensusState queries a consensus state associated with a client state at
// a given height.
func (c tendermintClient) GetConsensusState(chainName string, height uint64) (*tendermint.QueryConsensusStateResponse, error) {
	req := &tendermint.QueryConsensusStateRequest{
		ChainName:      chainName,
		RevisionHeight: height,
	}
	conn, err := c.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}

	res, err := tendermint.NewQueryClient(conn).ConsensusState(
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
func (c tendermintClient) ConsensusStates(chainName string) (*tendermint.QueryConsensusStatesResponse, error) {
	req := &tendermint.QueryConsensusStatesRequest{
		ChainName: chainName,
	}
	conn, err := c.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}

	res, err := tendermint.NewQueryClient(conn).ConsensusStates(
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
func (c tendermintClient) Relayers(chainName string) (*tendermint.QueryRelayersResponse, error) {
	req := &tendermint.QueryRelayersRequest{
		ChainName: chainName,
	}
	conn, err := c.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}

	res, err := tendermint.NewQueryClient(conn).Relayers(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, types.Wrap(err)
	}
	// todo ? change res to value?

	return res, nil
}
func (c tendermintClient) UpdateClient(msgUpdateClient tendermint.UpdateClientRequest) (*tendermint.MsgUpdateClientResponse, error) {
	req := &tendermint.MsgUpdateClient{}
	conn, err := c.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}
	res, err := tendermint.NewMsgClient(conn).UpdateClient(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, types.Wrap(err)
	}
	// todo ? change res to value?
	return res, nil
}
