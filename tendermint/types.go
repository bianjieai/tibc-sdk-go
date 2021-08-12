package tendermint

import (
	"fmt"
	sdk "github.com/irisnet/core-sdk-go/types"
	"sort"

	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
)

const (
	ModuleName = "tendermintClient"
)
var(
	_ tibctypes.ClientState = &ClientState{}
)

func (c QueryClientStatesResponse)Route() string  {
	return ModuleName
}
func (c QueryClientStatesResponse)Type() string  {
	return "queryTendermintClient"
}
func (c QueryClientStatesResponse)ValidateBasic() error  {
	return nil
}
func (c QueryClientStatesResponse)GetSignBytes() []byte{
	return nil
}
func (c QueryClientStatesResponse)GetSigners() []sdk.AccAddress{
	return nil
}

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

