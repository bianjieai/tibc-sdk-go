package client

import (
	"time"

	_go "github.com/confio/ics23/go"
	types "github.com/cosmos/cosmos-sdk/codec/types"
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
type Fraction struct {
	Numerator   uint64 `protobuf:"varint,1,opt,name=numerator,proto3" json:"numerator,omitempty"`
	Denominator uint64 `protobuf:"varint,2,opt,name=denominator,proto3" json:"denominator,omitempty"`
}

type State struct {
	ChainId                      string           `json:"chain_id,omitempty"`
	TrustLevel                   Fraction         `json:"trust_level"`
	TrustingPeriod               time.Duration    ` json:"trusting_period" `
	UnbondingPeriod              time.Duration    ` json:"unbonding_period" `
	MaxClockDrift                time.Duration    ` json:"max_clock_drift" `
	FrozenHeight                 Height           `json:"frozen_height" `
	LatestHeight                 Height           ` json:"latest_height"`
	ProofSpecs                   []*_go.ProofSpec `json:"proof_specs,omitempty" `
	UpgradePath                  []string         `json:"upgrade_path,omitempty" `
	AllowUpdateAfterExpiry       bool             `json:"allow_update_after_expiry,omitempty" `
	AllowUpdateAfterMisbehaviour bool             ` json:"allow_update_after_misbehaviour,omitempty" `
}

type UpdateClientRequest struct {
	ChainName string     `json:"chain_name,omitempty"`
	Header    *types.Any `json:"header,omitempty"`
	Signer    string     `json:"signer,omitempty"`
}
