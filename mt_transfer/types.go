package mt_transfer

import sdk "github.com/irisnet/core-sdk-go/types"

// Route Implements Msg
func (msg MsgMtTransfer) Route() string { return "MT" }

// Type Implements Msg
func (msg MsgMtTransfer) Type() string { return "tibc_mt_transfer" }

// GetSignBytes implements sdk.Msg.
func (msg MsgMtTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgMtTransfer) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

// ValidateBasic Implements Msg.
func (msg MsgMtTransfer) ValidateBasic() error {
	return nil
}
