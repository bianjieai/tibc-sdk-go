package packet

import "github.com/bianjieai/tibc-sdk-go/types"

const moduleName = "tibc" + "-" + "packet"

// TIBC packet sentinel errors
var (
	ErrSequenceSendNotFound     = types.Register(moduleName, 2, "sequence send not found")
	ErrSequenceReceiveNotFound  = types.Register(moduleName, 3, "sequence receive not found")
	ErrSequenceAckNotFound      = types.Register(moduleName, 4, "sequence acknowledgement not found")
	ErrInvalidPacket            = types.Register(moduleName, 5, "invalid packet")
	ErrInvalidAcknowledgement   = types.Register(moduleName, 6, "invalid acknowledgement")
	ErrPacketCommitmentNotFound = types.Register(moduleName, 7, "packet commitment not found")
	ErrPacketReceived           = types.Register(moduleName, 8, "packet already received")
	ErrAcknowledgementExists    = types.Register(moduleName, 9, "acknowledgement for packet already exists")
	ErrInvalidCleanPacket       = types.Register(moduleName, 10, "invalid clean packet")
	ErrInvalidProof             = types.Register(moduleName, 11, "invalid proof")
	ErrInvalidPrefix            = types.Register(moduleName, 12, "invalid prefix")
	ErrInvalidMerkleProof       = types.Register(moduleName, 13, "invalid merkle proof")
)
