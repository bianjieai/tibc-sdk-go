package types

import (
	"github.com/irisnet/core-sdk-go/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgNftTransfer{})
}
