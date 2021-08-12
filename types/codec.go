package types

import (
	"github.com/bianjieai/tibc-sdk-go/client"
	"github.com/irisnet/core-sdk-go/common/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)


func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface(
		"ibc.core.client.v1.ClientState",
		(*ClientState)(nil),
	)
	registry.RegisterInterface(
		"ibc.core.client.v1.ConsensusState",
		(*ConsensusState)(nil),
	)
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&client.MsgUpdateClient{},
	)
}