package tendermint

import (
	"github.com/bianjieai/tibc-sdk-go/modules/core/client"
	tibctypes "github.com/bianjieai/tibc-sdk-go/modules/types"
	"github.com/irisnet/core-sdk-go/common/codec"
	"github.com/irisnet/core-sdk-go/common/codec/types"
	cryptocodec "github.com/irisnet/core-sdk-go/common/crypto/codec"
	sdk "github.com/irisnet/core-sdk-go/types"
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
		&ConsensusState{},
	)
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&client.MsgUpdateClient{},
	)
}
