package tendermint

import (
	"github.com/bianjieai/tibc-sdk-go/client"
	sdk "github.com/irisnet/core-sdk-go/types"
)

type Client interface {
	sdk.Module

	QueryClientState(string) (ClientState, error)
	QueryConsensusState(chainName string, height uint64) (ConsensusState, error)

	UpdateClient(client.UpdateClientRequest) (*client.MsgUpdateClientResponse, error)
}
