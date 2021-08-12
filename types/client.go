package types

import 	"github.com/gogo/protobuf/proto"

// TendermintClientState defines the required common functions for light clients.
type ClientState interface {
	proto.Message

	//ClientType() string
	//GetLatestHeight() Height
	//Validate() error
	//GetDelayTime() uint64
	//GetDelayBlock() uint64
	//GetPrefix() Prefix

}

// TendermintConsensusState is the state of the consensus process
type ConsensusState interface {
	proto.Message

	//ClientType() string // Consensus kind
	//
	//// GetRoot returns the commitment root of the consensus state,
	//// which is used for key-value pair verification.
	//GetRoot() Root
	//
	//// GetTimestamp returns the timestamp (in nanoseconds) of the consensus state
	//GetTimestamp() uint64
	//
	//ValidateBasic() error
}
// Header is the consensus state update information
type Header interface {
	proto.Message

	//ClientType() string
	//GetHeight() Height
	//ValidateBasic() error
}