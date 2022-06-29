package nft_transfer

import sdk "github.com/irisnet/core-sdk-go/types"

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
