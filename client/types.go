package client

import (
	"fmt"
	"sort"

	"github.com/irisnet/core-sdk-go/client"
	"google.golang.org/grpc"
)

const (
	ModuleName = "client"
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

func (c *queryStateClient) NewGrpcConn(address string) {
	c.GRPCClient = client.NewGRPCClient(address)
}

func (c *queryStateClient) GenConn() (*grpc.ClientConn, error) {
	conn, err := c.GRPCClient.GenConn()
	return conn, err
}
