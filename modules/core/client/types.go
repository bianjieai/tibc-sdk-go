package client

import (
	"regexp"
	"sort"
	"strings"

	"github.com/bianjieai/tibc-sdk-go/types"
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
		return types.Wrapf(types.ErrInvalidAddress, "string could not be parsed as address: %v", err0)
	}
	header, err1 := types.UnpackHeader(m.Header)
	if err1 != nil {
		return err1
	}
	if err2 := header.ValidateBasic(); err2 != nil {
		return err2
	}
	return ClientIdentifierValidator(m.ChainName)
}
func ClientIdentifierValidator(id string) error {
	return defaultIdentifierValidator(id, 9, 64)
}

func defaultIdentifierValidator(id string, min, max int) error {
	if strings.TrimSpace(id) == "" {
		return types.Wrap(types.ErrInvalidID, "identifier cannot be blank")
	}
	// valid id MUST NOT contain "/" separator
	if strings.Contains(id, "/") {
		return types.Wrapf(types.ErrInvalidID, "identifier %s cannot contain separator '/'", id)
	}
	// valid id must fit the length requirements
	if len(id) < min || len(id) > max {
		return types.Wrapf(types.ErrInvalidID, "identifier %s has invalid length: %d, must be between %d-%d characters", id, len(id), min, max)
	}
	// valid id must contain only lower alphabetic characters
	if !IsValidID(id) {
		return types.Wrapf(
			types.ErrInvalidID,
			"identifier %s must contain only alphanumeric or the following characters: '.', '_', '+', '-', '#', '[', ']', '<', '>'",
			id,
		)
	}
	return nil
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
