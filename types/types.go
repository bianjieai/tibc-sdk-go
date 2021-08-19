package types

import (
	"errors"
	sdk "github.com/irisnet/core-sdk-go/types"

	coretypes "github.com/irisnet/core-sdk-go/common/codec/types"
)


// UnpackHeader unpacks an Any into a Header. It returns an error if the
// consensus state can't be unpacked into a Header.
func UnpackHeader(any *coretypes.Any) (Header, error) {
	if any == nil {
		return nil, errors.New("protobuf Any message cannot be nil")
	}

	header, ok := any.GetCachedValue().(Header)
	if !ok {
		return nil, errors.New("cannot unpack Any into Header")
	}
	return header, nil
}
// Route Implements Msg
func (msg MsgNftTransfer) Route() string { return "NFT" }

// Type Implements Msg
func (msg MsgNftTransfer) Type() string { return "tibc_nft_transfer" }

// GetSignBytes implements sdk.Msg.
func (msg MsgNftTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgNftTransfer) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

// ValidateBasic Implements Msg.
func (msg MsgNftTransfer) ValidateBasic() error {
	return nil
}