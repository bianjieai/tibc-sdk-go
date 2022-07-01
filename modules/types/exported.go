package types

import (
	"github.com/gogo/protobuf/proto"
)

// TODO:
type UpdateClientRequest struct {
	ChainName string `json:"chain_name"`
	// header to update the light client
	Header Header `json:"header"`
}

// ClientState defines the required common functions for light clients.
type ClientState interface {
	proto.Message

	ClientType() string
	GetLatestHeight() Height
	Validate() error
	GetDelayTime() uint64
	GetDelayBlock() uint64
	GetPrefix() Prefix
}

// ConsensusState is the state of the consensus process
type ConsensusState interface {
	proto.Message

	ClientType() string // Consensus kind

	// GetRoot returns the commitment root of the consensus state,
	// which is used for key-value pair verification.
	GetRoot() Root

	// GetTimestamp returns the timestamp (in nanoseconds) of the consensus state
	GetTimestamp() uint64

	ValidateBasic() error
}

// Header is the consensus state update information
type Header interface {
	proto.Message

	ClientType() string
	GetHeight() Height
	ValidateBasic() error
}

// Root implements spec:CommitmentRoot.
// A root is constructed from a set of key-value pairs,
// and the inclusion or non-inclusion of an arbitrary key-value pair
// can be proven with the proof.
type Root interface {
	GetHash() []byte
	Empty() bool
}

// Prefix implements spec:CommitmentPrefix.
// Prefix represents the common "prefix" that a set of keys shares.
type Prefix interface {
	Bytes() []byte
	Empty() bool
}

// Path implements spec:CommitmentPath.
// A path is the additional information provided to the verification function.
type Path interface {
	String() string
	Empty() bool
}

// Height is a wrapper interface over client.Height
// all clients must use the concrete implementation in types
type Height interface {
	IsZero() bool
	GetRevisionNumber() uint64
	GetRevisionHeight() uint64
	Increment() Height
	Decrement() (Height, bool)
	String() string
}

// PacketI defines the standard interface for IBC packets
type PacketI interface {
	GetSequence() uint64
	GetPort() string
	GetSourceChain() string
	GetDestChain() string
	GetRelayChain() string
	GetData() []byte
	ValidateBasic() error
}

// CleanPacketI defines the standard interface for TIBC clean packets
type CleanPacketI interface {
	GetSequence() uint64
	GetSourceChain() string
	GetDestChain() string
	GetRelayChain() string
	ValidateBasic() error
}
