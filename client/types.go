package client

import (
	"errors"
	"regexp"
	"sort"
	"strings"

	"github.com/bianjieai/tibc-sdk-go/types"
	coretypes "github.com/irisnet/core-sdk-go/common/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)

var IsValidID = regexp.MustCompile(`^[a-zA-Z0-9\.\_\+\-\#\[\]\<\>]+$`).MatchString

type IdentifiedClientStates []IdentifiedClientState
type ClientsConsensusStates []ClientConsensusStates

// Len implements sort.Interface
func (ics IdentifiedClientStates) Len() int { return len(ics) }

// Less implements sort.Interface
func (ics IdentifiedClientStates) Less(i, j int) bool { return ics[i].ChainName < ics[j].ChainName }

// Swap implements sort.Interface
func (ics IdentifiedClientStates) Swap(i, j int) { ics[i], ics[j] = ics[j], ics[i] }

// Sort is a helper function to sort the set of IdentifiedClientStates in place
func (ics IdentifiedClientStates) Sort() IdentifiedClientStates {
	sort.Sort(ics)
	return ics
}

func (m *MsgUpdateClient) Route() string {
	return "tibc"
}

func (m *MsgUpdateClient) Type() string {
	return "update_client"
}

func (m *MsgUpdateClient) ValidateBasic() error {
	_, err0 := sdk.AccAddressFromBech32(m.Signer)
	if err0 != nil {
		return errors.New("string could not be parsed as address")
	}
	header, err1 := UnpackHeader(m.Header)
	if err1 != nil {
		return err1
	}
	if err2 := header.ValidateBasic(); err2 != nil {
		return err2
	}
	return defaultIdentifierValidator(m.ChainName, 9, 64)
}

func defaultIdentifierValidator(id string, min, max int) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("identifier cannot be blank")
	}
	// valid id MUST NOT contain "/" separator
	if strings.Contains(id, "/") {
		return errors.New("identifier  cannot contain separator '/'")
	}
	// valid id must fit the length requirements
	if len(id) < min || len(id) > max {
		return errors.New("identifier" + id + "  has invalid length:" + string(len(id)) + ", must be between " + string(min) + "- " + string(max) + " characters")
	}
	// valid id must contain only lower alphabetic characters
	if !IsValidID(id) {
		return errors.New("identifier " + id + " must contain only alphanumeric or the following characters: '.', '_', '+', '-', '#', '[', ']', '<', '>'")
	}
	return nil
}

// UnpackHeader unpacks an Any into a Header. It returns an error if the
// consensus state can't be unpacked into a Header.
func UnpackHeader(any *coretypes.Any) (types.Header, error) {
	if any == nil {
		return nil, errors.New("protobuf Any message cannot be nil")
	}

	header, ok := any.GetCachedValue().(types.Header)
	if !ok {
		return nil, errors.New("cannot unpack Any into Header")
	}
	return header, nil
}

func (m *MsgUpdateClient) GetSignBytes() []byte {
	return []byte(m.Signer)
}

func (m *MsgUpdateClient) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(m.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}
