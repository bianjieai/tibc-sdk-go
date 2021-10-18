package tendermint

import (
	"github.com/bianjieai/tibc-sdk-go/client"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	"github.com/irisnet/core-sdk-go/codec"
	"github.com/irisnet/core-sdk-go/codec/types"
	cryptocodec "github.com/irisnet/core-sdk-go/crypto/codec"
	sdk "github.com/irisnet/core-sdk-go/types"
)

var (
	amino = codec.NewLegacyAmino()
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
		&ConsensusState{},
	)
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&client.MsgUpdateClient{},
	)
}
