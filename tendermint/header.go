package tendermint

import (
	"bytes"
	"errors"
	"github.com/bianjieai/tibc-sdk-go/client"
	"time"

	commitmenttypes "github.com/bianjieai/tibc-sdk-go/commitment"
	"github.com/bianjieai/tibc-sdk-go/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

var _ types.Header = &Header{}

// ConsensusState returns the updated consensus state associated with the header
func (h Header) ConsensusState() *ConsensusState {
	return &ConsensusState{
		Timestamp:          h.GetTime(),
		Root:               commitmenttypes.NewMerkleRoot(h.Header.GetAppHash()),
		NextValidatorsHash: h.Header.NextValidatorsHash,
	}
}

// ClientType defines that the Header is a Tendermint consensus algorithm
func (h Header) ClientType() string {
	return "007-tendermint"
}

// GetHeight returns the current height. It returns 0 if the tendermint
// header is nil.
// NOTE: the header.Header is checked to be non nil in ValidateBasic.
func (h Header) GetHeight() types.Height {
	revision := client.ParseChainID(h.Header.ChainID)
	return client.NewHeight(revision, uint64(h.Header.Height))
}

// GetTime returns the current block timestamp. It returns a zero time if
// the tendermint header is nil.
// NOTE: the header.Header is checked to be non nil in ValidateBasic.
func (h Header) GetTime() time.Time {
	return h.Header.Time
}

// ValidateBasic calls the SignedHeader ValidateBasic function and checks
// that validatorsets are not nil.
// NOTE: TrustedHeight and TrustedValidators may be empty when creating client
// with MsgCreateClient
func (h Header) ValidateBasic() error {
	if h.SignedHeader == nil {
		return errors.New( "tendermint signed header cannot be nil")
	}
	if h.Header == nil {
		return errors.New( "tendermint header cannot be nil")
	}
	tmSignedHeader, err := tmtypes.SignedHeaderFromProto(h.SignedHeader)
	if err != nil {
		return errors.New( "header is not a tendermint header")
	}
	if err := tmSignedHeader.ValidateBasic(h.Header.GetChainID()); err != nil {
		return errors.New( "header failed basic validation")
	}

	// TrustedHeight is less than Header for updates
	// and less than or equal to Header for misbehaviour
	if h.TrustedHeight.GT(h.GetHeight()) {
		return errors.New( "TrustedHeight must be less than or equal to header height ")
	}

	if h.ValidatorSet == nil {
		return errors.New( "validator set is nil")
	}
	tmValset, err := tmtypes.ValidatorSetFromProto(h.ValidatorSet)
	if err != nil {
		return errors.New( "validator set is not tendermint validator set")
	}
	if !bytes.Equal(h.Header.ValidatorsHash, tmValset.Hash()) {
		return errors.New( "validator set does not match hash")
	}
	return nil
}
