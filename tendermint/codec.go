package tendermint

import (
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	"github.com/irisnet/core-sdk-go/common/codec"
	"github.com/irisnet/core-sdk-go/common/codec/types"
	cryptocodec "github.com/irisnet/core-sdk-go/common/crypto/codec"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*tibctypes.ClientState)(nil),
		&ClientState{},
	)
	registry.RegisterImplementations(
		(*tibctypes.ConsensusState)(nil),
		&ClientConsensusStates{},
	)
	registry.RegisterInterface(
		"tendermint.ClientState",
		(*tibctypes.ClientState)(nil),
	)
	registry.RegisterInterface(
		"ibc.core.client.v1.ConsensusState",
		(*tibctypes.ConsensusState)(nil),
	)



	//registry.RegisterImplementations(
	//	(*sdk.Msg)(nil),
	//	&MsgUpdateClient{},
	//	)
}
