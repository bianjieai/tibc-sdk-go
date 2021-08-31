package tendermint

import (
	"strings"

	"github.com/bianjieai/tibc-sdk-go/types"
	tmmath "github.com/tendermint/tendermint/libs/math"
	"github.com/tendermint/tendermint/light"
)

func (cs ClientState) GetLatestHeight() types.Height {
	return cs.LatestHeight
}

// GetChainId returns the chain-id
func (cs ClientState) GetChainId() string {
	return cs.ChainId
}

// ClientType is tendermint.
func (cs ClientState) ClientType() string {
	return "007-tendermint"
}

// GetDelayBlock returns the number of blocks delayed in transaction confirmation.
func (cs ClientState) GetDelayBlock() uint64 {
	return 0
}

// GetDelayTime returns the period of transaction confirmation delay.
func (cs ClientState) GetDelayTime() uint64 {
	return cs.TimeDelay
}

// GetPrefix returns the prefix path for proof key.
func (cs ClientState) GetPrefix() types.Prefix {
	return &cs.MerklePrefix
}

// Validate performs a basic validation of the client state fields.
func (cs ClientState) Validate() error {
	if strings.TrimSpace(cs.ChainId) == "" {
		return types.Wrap(ErrInvalidChainID, "chain id cannot be empty string")
	}
	if err := light.ValidateTrustLevel(cs.TrustLevel.ToTendermint()); err != nil {
		return err
	}
	if cs.TrustingPeriod == 0 {
		return types.Wrap(ErrInvalidTrustingPeriod, "trusting period cannot be zero")
	}
	if cs.UnbondingPeriod == 0 {
		return types.Wrap(ErrInvalidUnbondingPeriod, "unbonding period cannot be zero")
	}
	if cs.MaxClockDrift == 0 {
		return types.Wrap(ErrInvalidMaxClockDrift, "max clock drift cannot be zero")
	}
	if cs.LatestHeight.RevisionHeight == 0 {
		return types.Wrapf(ErrInvalidHeaderHeight, "tendermint revision height cannot be zero")
	}
	if cs.TrustingPeriod >= cs.UnbondingPeriod {
		return types.Wrapf(
			ErrInvalidTrustingPeriod,
			"trusting period (%s) should be < unbonding period (%s)", cs.TrustingPeriod, cs.UnbondingPeriod,
		)
	}
	if cs.ProofSpecs == nil {
		return types.Wrap(ErrInvalidProofSpecs, "proof specs cannot be nil for tm client")
	}
	for i, spec := range cs.ProofSpecs {
		if spec == nil {
			return types.Wrapf(ErrInvalidProofSpecs, "proof spec cannot be nil at index: %d", i)
		}
	}
	return nil
}

// ToTendermint converts Fraction to tmmath.Fraction
func (f Fraction) ToTendermint() tmmath.Fraction {
	return tmmath.Fraction{
		Numerator:   f.Numerator,
		Denominator: f.Denominator,
	}
}
