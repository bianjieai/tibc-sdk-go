package client

import (
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)


type ChainClient interface {
	sdk.Module

	GetClientState(chainName string) (tibctypes.ClientState, error)
	GetClientStates() ([]tibctypes.ClientState, error)
	GetConsensusState(chainName string, height uint64) (tibctypes.ConsensusState, error)
	GetConsensusStates(chainName string) ([]tibctypes.ConsensusState, error)
	Relayers(chainName string) ([]string, error)
	UpdateClient( tibctypes.UpdateClientRequest,  sdk.BaseTx)(sdk.ResultTx, sdk.Error)
}



