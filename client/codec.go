package client

import (
	"github.com/irisnet/core-sdk-go/common/codec"
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

//
//func RegisterInterfaces(registry types.InterfaceRegistry) {
//	registry.RegisterImplementations(
//		(*sdk.Msg)(nil),
//		& {},
//	)
//}
