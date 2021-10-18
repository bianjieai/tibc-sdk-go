package packet

import (
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
		(*sdk.Msg)(nil),
		&MsgAcknowledgement{},
		&MsgCleanPacket{},
		&MsgRecvPacket{},
		&MsgRecvCleanPacket{},
	)
}
