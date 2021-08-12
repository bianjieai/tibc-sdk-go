package tendermint

import (
	"github.com/irisnet/core-sdk-go/common/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)


type ChainClient interface {
	sdk.Module

	GetClientState(chainName string) (*QueryClientStateResponse, error)
	GetClientStates() (*QueryClientStatesResponse, error)
	GetConsensusState(chainName string, height uint64) (*QueryConsensusStateResponse, error)
	ConsensusStates(chainName string) (*QueryConsensusStatesResponse, error)
	Relayers(chainName string) (*QueryRelayersResponse, error)

}
type TendermintFraction struct {
	Numerator   uint64 `protobuf:"varint,1,opt,name=numerator,proto3" json:"numerator,omitempty"`
	Denominator uint64 `protobuf:"varint,2,opt,name=denominator,proto3" json:"denominator,omitempty"`
}



type UpdateClientRequest struct {
	ChainName string     `json:"chain_name,omitempty"`
	Header    *types.Any `json:"header,omitempty"`
	Signer    string     `json:"signer,omitempty"`
}

type Prefix interface {
	Bytes() []byte
	Empty() bool
}
type Root interface {
	GetHash() []byte
	Empty() bool
}
