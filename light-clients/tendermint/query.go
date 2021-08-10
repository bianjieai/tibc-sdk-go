package tendermint

import (
	"encoding/json"

	"github.com/bianjieai/tibc-sdk-go/client"
	"github.com/irisnet/core-sdk-go/common/codec"
	sdk "github.com/irisnet/core-sdk-go/types"
)

type tendermintLightClient struct {
	cli client.Client
}

func NewClient(bc sdk.BaseClient, cdc codec.Marshaler) (*tendermintLightClient, error) {
	return &tendermintLightClient{
		cli: client.NewClient(bc, cdc),
	}, nil
}

func (q *tendermintLightClient) QueryClientState(chainName string) (ClientState, error) {
	responseClientState, err := q.cli.GetClientState(chainName)
	if err != nil {
		return ClientState{}, sdk.Wrap(err)
	}
	res := ClientState{}
	err = json.Unmarshal(responseClientState.ClientState.Value, &res)
	if err != nil {
		return ClientState{}, sdk.Wrap(err)
	}
	return res, nil
}

func (q *tendermintLightClient) QueryConsensusState(chainName string, height uint64) (ConsensusState, error) {
	responseConsensusState, err := q.cli.GetConsensusState(chainName, height)
	if err != nil {
		return ConsensusState{}, sdk.Wrap(err)
	}
	res := ConsensusState{}
	err = json.Unmarshal(responseConsensusState.ConsensusState.Value, &res)
	if err != nil {
		return ConsensusState{}, sdk.Wrap(err)
	}
	return res, nil
}
