package types

import (
 	"github.com/irisnet/core-sdk-go/common/codec/types"
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

}