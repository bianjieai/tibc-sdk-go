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
		&ConsensusState{},
	)
}
//func UnpackClientState(any *types.Any) (ClientState, error) {
//	if any == nil {
//		return nil, errors.New("protobuf Any message cannot be nil")
//	}
//
//	clientState, ok := any.GetCachedValue().(exported.ClientState)
//	if !ok {
//		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnpackAny, "cannot unpack Any into ClientState %T", any)
//	}
//
//	return clientState, nil
//}
