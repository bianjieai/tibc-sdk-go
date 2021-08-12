package client

import (
	"fmt"
	sdk "github.com/irisnet/core-sdk-go/types"
	"sort"
)

const (
	ModuleName = "tendermintClient"
)


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

// String returns a string representation of Height
func (h Height) String() string {
	return fmt.Sprintf("%d-%d", h.RevisionNumber, h.RevisionHeight)
}


func (m *MsgUpdateClient) Route() string {
	return "tibc"
}

func (m *MsgUpdateClient) Type() string {
	return "update_client"
}

func (m *MsgUpdateClient) ValidateBasic() error {
	return nil
}

func (m *MsgUpdateClient) GetSignBytes() []byte {
	panic("IBC messages do not support amino")}

func (m *MsgUpdateClient) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(m.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}}
