package packet

import (
	"github.com/bianjieai/tibc-sdk-go/modules/core/commitment"
	tibctypes "github.com/bianjieai/tibc-sdk-go/modules/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)

var _ sdk.Msg = &MsgRecvPacket{}

// Route implements sdk.Msg
func (msg MsgRecvPacket) Route() string {
	return "tibc"
}

// GetSignBytes implements sdk.Msg. The function will panic since it is used
// for amino transaction verification which TIBC does not support.
func (msg MsgRecvPacket) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgRecvPacket) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

// ValidateBasic implements sdk.Msg
func (msg MsgRecvPacket) ValidateBasic() error {
	if len(msg.ProofCommitment) == 0 {
		return tibctypes.Wrap(commitment.ErrInvalidProof, "cannot submit an empty proof")
	}
	if msg.ProofHeight.IsZero() {
		return tibctypes.Wrap(tibctypes.ErrInvalidHeight, "proof height must be non-zero")
	}
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return tibctypes.Wrapf(tibctypes.ErrInvalidAddress, "string could not be parsed as address: %v", err)
	}
	return msg.Packet.ValidateBasic()
}

// Type implements sdk.Msg
func (msg MsgRecvPacket) Type() string {
	return "recv_packet"
}

var _ tibctypes.PacketI = (*Packet)(nil)

// NewPacket creates a new Packet instance. It panics if the provided
// packet data interface is not registered.
func NewPacket(
	data []byte,
	sequence uint64, sourceChain, destinationChain, relayChain,
	port string,
) Packet {
	return Packet{
		Data:             data,
		Sequence:         sequence,
		SourceChain:      sourceChain,
		DestinationChain: destinationChain,
		RelayChain:       relayChain,
		Port:             port,
	}
}

// GetSequence implements PacketI interface
func (p Packet) GetSequence() uint64 { return p.Sequence }

// GetPort implements PacketI interface
func (p Packet) GetPort() string { return p.Port }

// GetSourceChain implements PacketI interface
func (p Packet) GetSourceChain() string { return p.SourceChain }

// GetDestinationChain implements PacketI interface
func (p Packet) GetDestChain() string { return p.DestinationChain }

// GetRelayChain implements PacketI interface
func (p Packet) GetRelayChain() string { return p.RelayChain }

// GetData implements PacketI interface
func (p Packet) GetData() []byte { return p.Data }

// ValidateBasic implements PacketI interface
func (p Packet) ValidateBasic() error {
	if p.Sequence == 0 {
		return tibctypes.Wrap(ErrInvalidPacket, "packet sequence cannot be 0")
	}
	if len(p.Data) == 0 {
		return tibctypes.Wrap(ErrInvalidPacket, "packet data bytes cannot be empty")
	}
	return nil
}

var _ sdk.Msg = &MsgAcknowledgement{}

// Route implements sdk.Msg
func (msg MsgAcknowledgement) Route() string {
	return "tibc"
}

// ValidateBasic implements sdk.Msg
func (msg MsgAcknowledgement) ValidateBasic() error {
	if len(msg.ProofAcked) == 0 {
		return tibctypes.Wrap(commitment.ErrInvalidProof, "cannot submit an empty proof")
	}
	if msg.ProofHeight.IsZero() {
		return tibctypes.Wrap(tibctypes.ErrInvalidHeight, "proof height must be non-zero")
	}
	if len(msg.Acknowledgement) == 0 {
		return tibctypes.Wrap(ErrInvalidAcknowledgement, "ack bytes cannot be empty")
	}
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return tibctypes.Wrapf(tibctypes.ErrInvalidAddress, "string could not be parsed as address: %v", err)
	}
	return msg.Packet.ValidateBasic()
}

// GetSignBytes implements sdk.Msg. The function will panic since it is used
// for amino transaction verification which TIBC does not support.
func (msg MsgAcknowledgement) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgAcknowledgement) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

// Type implements sdk.Msg
func (msg MsgAcknowledgement) Type() string {
	return "acknowledge_packet"
}

var _ sdk.Msg = &MsgCleanPacket{}

// Route implements sdk.Msg
func (msg MsgCleanPacket) Route() string {
	return "tibc"
}

// ValidateBasic implements sdk.Msg
func (msg MsgCleanPacket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return tibctypes.Wrapf(tibctypes.ErrInvalidAddress, "string could not be parsed as address: %v", err)
	}
	return msg.CleanPacket.ValidateBasic()
}

// GetSignBytes implements sdk.Msg. The function will panic since it is used
// for amino transaction verification which TIBC does not support.
func (msg MsgCleanPacket) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgCleanPacket) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

// Type implements sdk.Msg
func (msg MsgCleanPacket) Type() string {
	return "clean_packet"
}

var _ sdk.Msg = &MsgRecvCleanPacket{}

// Route implements sdk.Msg
func (msg MsgRecvCleanPacket) Route() string {
	return "tibc"
}

// ValidateBasic implements sdk.Msg
func (msg MsgRecvCleanPacket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return tibctypes.Wrapf(tibctypes.ErrInvalidAddress, "string could not be parsed as address: %v", err)
	}
	return msg.CleanPacket.ValidateBasic()
}

// GetSignBytes implements sdk.Msg. The function will panic since it is used
// for amino transaction verification which TIBC does not support.
func (msg MsgRecvCleanPacket) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgRecvCleanPacket) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

// Type implements sdk.Msg
func (msg MsgRecvCleanPacket) Type() string {
	return "recv_clean_packet"
}

var _ tibctypes.CleanPacketI = (*CleanPacket)(nil)

// GetSequence implements PacketI interface
func (p CleanPacket) GetSequence() uint64 { return p.Sequence }

// GetSourceChain implements PacketI interface
func (p CleanPacket) GetSourceChain() string { return p.SourceChain }

// GetDestinationChain implements PacketI interface
func (p CleanPacket) GetDestChain() string { return p.DestinationChain }

// GetRelayChain implements PacketI interface
func (p CleanPacket) GetRelayChain() string { return p.RelayChain }

// ValidateBasic implements PacketI interface
func (p CleanPacket) ValidateBasic() error {
	if p.Sequence == 0 {
		return tibctypes.Wrap(ErrInvalidPacket, "packet sequence cannot be 0")
	}
	return nil
}
