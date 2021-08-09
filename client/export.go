package client

import (
	"github.com/gogo/protobuf/proto"

	"github.com/bianjieai/tibc-sdk-go/commitment"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// TypeClientMisbehaviour is the shared evidence misbehaviour type
	TypeClientMisbehaviour string = "client_misbehaviour"

	// Tendermint is used to indicate that the client uses the Tendermint Consensus Algorithm.
	Tendermint string = "007-tendermint"

	// BSC is the client type for a bianance smart chain client.
	BSC string = "008-bsc"

	// Fabric is the client type for a hyperledge fabric client.
	Fabric string = "009-fabric"
)

// ClientState defines the required common functions for light clients.
type ClientState interface {
	proto.Message

	ClientType() string
	GetLatestHeight() Height
	Validate() error
	DelayTime() uint64
	DelayBlock() uint64
	Prefix() commitment.Prefix

	// Initialization function
	// Clients must validate the initial consensus state, and may store any client-specific metadata
	// necessary for correct light client operation
	Initialize(sdk.Context, codec.BinaryMarshaler, sdk.KVStore, ConsensusState) error

	// Genesis function
	ExportMetadata(sdk.KVStore) []GenesisMetadata

	// Update and Misbehaviour functions

	CheckHeaderAndUpdateState(sdk.Context, codec.BinaryMarshaler, sdk.KVStore, Header) (ClientState, ConsensusState, error)

	// State verification functions

	VerifyPacketCommitment(
		store sdk.KVStore,
		cdc codec.BinaryMarshaler,
		height Height,
		proof []byte,
		sourceChain,
		destChain string,
		sequence uint64,
		commitmentBytes []byte,
	) error

	VerifyPacketAcknowledgement(
		store sdk.KVStore,
		cdc codec.BinaryMarshaler,
		height Height,
		proof []byte,
		sourceChain,
		destChain string,
		sequence uint64,
		acknowledgement []byte,
	) error

	VerifyPacketCleanCommitment(
		store sdk.KVStore,
		cdc codec.BinaryMarshaler,
		height Height,
		proof []byte,
		sourceChain string,
		destChain string,
		sequence uint64,
		cleanCommitmentBytes []byte,
	) error
}

// ConsensusState is the state of the consensus process
type ConsensusState interface {
	proto.Message

	ClientType() string // Consensus kind

	// GetRoot returns the commitment root of the consensus state,
	// which is used for key-value pair verification.
	GetRoot() commitment.Root

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
