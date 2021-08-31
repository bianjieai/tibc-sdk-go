package tendermint

import (
	"time"

	"github.com/bianjieai/tibc-sdk-go/commitment"
	"github.com/bianjieai/tibc-sdk-go/types"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	tmtypes "github.com/tendermint/tendermint/types"
)

// NewConsensusState creates a new ConsensusState instance.
func NewConsensusState(
	timestamp time.Time, root commitment.MerkleRoot, nextValsHash tmbytes.HexBytes,
) *ConsensusState {
	return &ConsensusState{
		Timestamp:          timestamp,
		Root:               root,
		NextValidatorsHash: nextValsHash,
	}
}

// ClientType returns Tendermint
func (ConsensusState) ClientType() string {
	return "007-tendermint"
}

// GetRoot returns the commitment Root for the specific
func (cs ConsensusState) GetRoot() types.Root {
	return cs.Root
}

// GetTimestamp returns block time in nanoseconds of the header that created consensus state
func (cs ConsensusState) GetTimestamp() uint64 {
	return uint64(cs.Timestamp.UnixNano())
}

// ValidateBasic defines a basic validation for the tendermint consensus state.
// NOTE: ProcessedTimestamp may be zero if this is an initial consensus state passed in by relayer
// as opposed to a consensus state constructed by the chain.
func (cs ConsensusState) ValidateBasic() error {
	if cs.Root.Empty() {
		return types.Wrap(types.ErrInvalidConsensus, "root cannot be empty")
	}
	if err := tmtypes.ValidateHash(cs.NextValidatorsHash); err != nil {
		return types.Wrap(err, "next validators hash is invalid")
	}
	if cs.Timestamp.Unix() <= 0 {
		return types.Wrap(types.ErrInvalidConsensus, "timestamp must be a positive Unix time")
	}
	return nil
}
