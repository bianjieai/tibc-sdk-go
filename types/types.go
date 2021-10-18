package types

import (
	coretypes "github.com/irisnet/core-sdk-go/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)

// UnpackHeader unpacks an Any into a Header. It returns an error if the
// consensus state can't be unpacked into a Header.
func UnpackHeader(any *coretypes.Any) (Header, error) {
	if any == nil {
		return nil, Wrap(ErrUnpackAny, "protobuf Any message cannot be nil")
	}

	header, ok := any.GetCachedValue().(Header)
	if !ok {
		return nil, Wrapf(ErrUnpackAny, "cannot unpack Any into Header %T", any)
	}

	return header, nil
}

// Route Implements Msg
func (msg MsgNftTransfer) Route() string { return "NFT" }

// Type Implements Msg
func (msg MsgNftTransfer) Type() string { return "tibc_nft_transfer" }

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
