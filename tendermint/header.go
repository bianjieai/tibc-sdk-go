package tendermint

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bianjieai/tibc-sdk-go/client"
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
	if !client.IsRevisionFormat(h.Header.ChainID) {
		// chainID is not in revision format, return 0 as default
		return client.NewHeight(0, uint64(h.Header.Height))
	} else {
		splitStr := strings.Split(h.Header.ChainID, "-")
		revision, err := strconv.ParseUint(splitStr[len(splitStr)-1], 10, 64)
		// sanity check: error should always be nil since regex only allows numbers in last element
		if err != nil {
			panic(fmt.Sprintf("regex allowed non-number value as last split element for chainID: %s", h.Header.ChainID))
		}
		return client.NewHeight(revision, uint64(h.Header.Height))
	}

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
		return types.Wrap(types.ErrInvalidHeader, "tendermint signed header cannot be nil")
	}
	if h.Header == nil {
		return types.Wrap(types.ErrInvalidHeader, "tendermint header cannot be nil")
	}
	tmSignedHeader, err := tmtypes.SignedHeaderFromProto(h.SignedHeader)
	if err != nil {
		return types.Wrap(err, "header is not a tendermint header")
	}
	if err := tmSignedHeader.ValidateBasic(h.Header.GetChainID()); err != nil {
		return types.Wrap(err, "header failed basic validation")
	}

	if h.ValidatorSet == nil {
		return types.Wrap(types.ErrInvalidHeader, "validator set is nil")
	}
	tmValset, err := tmtypes.ValidatorSetFromProto(h.ValidatorSet)
	if err != nil {
		return types.Wrap(err, "validator set is not tendermint validator set")
	}
	if !bytes.Equal(h.Header.ValidatorsHash, tmValset.Hash()) {
		return types.Wrap(types.ErrInvalidHeader, "validator set does not match hash")
	}
	return nil
}
