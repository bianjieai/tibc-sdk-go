package types

type UpdateClientRequest struct {
	ChainName string `json:"chain_name"`
	// header to update the light client
	Header Header `json:"header"`

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

// Prefix implements spec:CommitmentPrefix.
// Prefix represents the common "prefix" that a set of keys shares.
type Prefix interface {
	Bytes() []byte
	Empty() bool
}

// Root implements spec:CommitmentRoot.
// A root is constructed from a set of key-value pairs,
// and the inclusion or non-inclusion of an arbitrary key-value pair
// can be proven with the proof.
type Root interface {
	GetHash() []byte
	Empty() bool
}

// Path implements spec:CommitmentPath.
// A path is the additional information provided to the verification function.
type Path interface {
	String() string
	Empty() bool
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

