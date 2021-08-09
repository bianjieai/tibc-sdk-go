package client

import (
	"time"

	_go "github.com/confio/ics23/go"
	"github.com/irisnet/core-sdk-go/types"
)

type ChainClient interface {
	types.Module

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
	ChainId    string   `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	TrustLevel Fraction `protobuf:"bytes,2,opt,name=trust_level,json=trustLevel,proto3" json:"trust_level" yaml:"trust_level"`
	// duration of the period since the LastestTimestamp during which the
	// submitted headers are valid for upgrade
	TrustingPeriod time.Duration `protobuf:"bytes,3,opt,name=trusting_period,json=trustingPeriod,proto3,stdduration" json:"trusting_period" yaml:"trusting_period"`
	// duration of the staking unbonding period
	UnbondingPeriod time.Duration `protobuf:"bytes,4,opt,name=unbonding_period,json=unbondingPeriod,proto3,stdduration" json:"unbonding_period" yaml:"unbonding_period"`
	// defines how much new (untrusted) header's Time can drift into the future.
	MaxClockDrift time.Duration `protobuf:"bytes,5,opt,name=max_clock_drift,json=maxClockDrift,proto3,stdduration" json:"max_clock_drift" yaml:"max_clock_drift"`
	// Block height when the client was frozen due to a misbehaviour
	FrozenHeight Height `protobuf:"bytes,6,opt,name=frozen_height,json=frozenHeight,proto3" json:"frozen_height" yaml:"frozen_height"`
	// Latest height the client was updated to
	LatestHeight Height `protobuf:"bytes,7,opt,name=latest_height,json=latestHeight,proto3" json:"latest_height" yaml:"latest_height"`
	// Proof specifications used in verifying counterparty state
	ProofSpecs []*_go.ProofSpec `protobuf:"bytes,8,rep,name=proof_specs,json=proofSpecs,proto3" json:"proof_specs,omitempty" yaml:"proof_specs"`
	// Path at which next upgraded client will be committed.
	// Each element corresponds to the key for a single CommitmentProof in the chained proof.
	// NOTE: ClientState must stored under `{upgradePath}/{upgradeHeight}/clientState`
	// ConsensusState must be stored under `{upgradepath}/{upgradeHeight}/consensusState`
	// For SDK chains using the default upgrade module, upgrade_path should be []string{"upgrade", "upgradedIBCState"}`
	UpgradePath []string `protobuf:"bytes,9,rep,name=upgrade_path,json=upgradePath,proto3" json:"upgrade_path,omitempty" yaml:"upgrade_path"`
	// This flag, when set to true, will allow governance to recover a client
	// which has expired
	AllowUpdateAfterExpiry bool `protobuf:"varint,10,opt,name=allow_update_after_expiry,json=allowUpdateAfterExpiry,proto3" json:"allow_update_after_expiry,omitempty" yaml:"allow_update_after_expiry"`
	// This flag, when set to true, will allow governance to unfreeze a client
	// whose chain has experienced a misbehaviour event
	AllowUpdateAfterMisbehaviour bool `protobuf:"varint,11,opt,name=allow_update_after_misbehaviour,json=allowUpdateAfterMisbehaviour,proto3" json:"allow_update_after_misbehaviour,omitempty" yaml:"allow_update_after_misbehaviour"`
}
