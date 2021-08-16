package tibc_sdk_go

import (
	"context"
	"github.com/bianjieai/tibc-sdk-go/client"
	"github.com/bianjieai/tibc-sdk-go/tendermint"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	commoncodec "github.com/irisnet/core-sdk-go/common/codec"
	cryptotypes "github.com/irisnet/core-sdk-go/common/codec/types"
	"github.com/irisnet/core-sdk-go/types"
)

type Client struct {
	types.BaseClient
	commoncodec.Marshaler
}

func NewClient(bc types.BaseClient, cdc commoncodec.Marshaler) Client {
	return Client{
		BaseClient: bc,
		Marshaler:  cdc,
	}
}

func (c Client) RegisterInterfaceTypes(registry cryptotypes.InterfaceRegistry) {
	tendermint.RegisterInterfaces(registry)
}

func (c Client) Name() string {
	return "tibc"
}

// GetClientState queries an IBC light client.
func (c Client) GetClientState(chainName string) (tibctypes.ClientState, error) {
	var clientState tibctypes.ClientState
	in := &client.QueryClientStateRequest{
		ChainName: chainName,
	}

	conn, err := c.GenConn()
	if err != nil {
		return clientState, types.Wrap(err)
	}

	res, err := client.NewQueryClient(conn).ClientState(
		context.Background(),
		in,
	)
	if err != nil {
		return clientState, types.Wrap(err)
	}

	if err := c.Marshaler.UnpackAny(res.ClientState, &clientState); err != nil {
		return clientState, types.Wrap(err)
	}
	return clientState, nil

}

// GetClientStates queries all the IBC light clients of a chain.
func (c Client) GetClientStates() ([]tibctypes.ClientState, error) {
	in := &client.QueryClientStatesRequest{}
	conn, err := c.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}
	res, err := client.NewQueryClient(conn).ClientStates(
		context.Background(),
		in,
	)
	if err != nil {
		return nil, types.Wrap(err)
	}
	clientState := make([]tibctypes.ClientState, len(res.ClientStates))
	for index, value := range res.ClientStates {
		if err := c.Marshaler.UnpackAny(value.ClientState, &clientState[index]); err != nil {
			return nil, types.Wrap(err)
		}
	}

	return clientState, err
}

// GetConsensusState queries a consensus state associated with a client state at
// a given height.
func (c Client) GetConsensusState(chainName string, height uint64) (tibctypes.ConsensusState, error) {
	req := &client.QueryConsensusStateRequest{
		ChainName:      chainName,
		RevisionHeight: height,
	}

	conn, err := c.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}

	res, err := client.NewQueryClient(conn).ConsensusState(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, types.Wrap(err)
	}
	var consensusState tibctypes.ConsensusState

	if err := c.Marshaler.UnpackAny(res.ConsensusState, &consensusState); err != nil {
		return nil, types.Wrap(err)
	}

	return consensusState, nil
}

// GetConsensusStates queries all the consensus state associated with a given
// client.
func (c Client) GetConsensusStates(chainName string) ([]tibctypes.ConsensusState, error) {
	req := &client.QueryConsensusStatesRequest{
		ChainName: chainName,
	}
	conn, err := c.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}

	res, err := client.NewQueryClient(conn).ConsensusStates(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, types.Wrap(err)
	}
	ConsensusState := make([]tibctypes.ConsensusState, len(res.ConsensusStates))
	for index, value := range res.ConsensusStates {
		if err := c.Marshaler.UnpackAny(value.ConsensusState, &ConsensusState[index]); err != nil {
			return nil, types.Wrap(err)
		}
	}
	return ConsensusState, nil
}

// Relayers queries all the relayers associated with a given
// client.
func (c Client) Relayers(chainName string) ([]string, error) {
	req := &client.QueryRelayersRequest{
		ChainName: chainName,
	}
	conn, err := c.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}

	relay, err := client.NewQueryClient(conn).Relayers(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, types.Wrap(err)
	}

	return relay.Relayers, nil
}

func (c Client) UpdateClient(req tibctypes.UpdateClientRequest, baseTx types.BaseTx) (types.ResultTx, types.Error) {
	owner, err := c.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return types.ResultTx{}, types.Wrap(err)
	}
	res, errs := cryptotypes.NewAnyWithValue(req.Header)
	if errs != nil {
		return types.ResultTx{}, types.Wrap(errs)
	}
	msg := &client.MsgUpdateClient{
		ChainName: req.ChainName,
		// header to update the light client
		Header: res,
		// signer address
		Signer: owner.String(),
	}
	return c.BuildAndSend([]types.Msg{msg}, baseTx)
}
