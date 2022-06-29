package tendermint

import (
	"github.com/bianjieai/tibc-sdk-go/types"
)

const (
	SubModuleName = "tendermint-client"
	moduleName    = "tibc" + "-" + SubModuleName
)

// IBC tendermint client sentinel errors
var (
	ErrInvalidChainID         = types.Register(moduleName, 2, "invalid chain-id")
	ErrInvalidTrustingPeriod  = types.Register(moduleName, 3, "invalid trusting period")
	ErrInvalidUnbondingPeriod = types.Register(moduleName, 4, "invalid unbonding period")
	ErrInvalidHeaderHeight    = types.Register(moduleName, 5, "invalid header height")
	ErrInvalidHeader          = types.Register(moduleName, 6, "invalid header")
	ErrInvalidMaxClockDrift   = types.Register(moduleName, 7, "invalid max clock drift")
	ErrProcessedTimeNotFound  = types.Register(moduleName, 8, "processed time not found")
	ErrDelayPeriodNotPassed   = types.Register(moduleName, 9, "packet-specified delay period has not been reached")
	ErrTrustingPeriodExpired  = types.Register(moduleName, 10, "time since latest trusted state has passed the trusting period")
	ErrUnbondingPeriodExpired = types.Register(moduleName, 11, "time since latest trusted state has passed the unbonding period")
	ErrInvalidProofSpecs      = types.Register(moduleName, 12, "invalid proof specs")
	ErrInvalidValidatorSet    = types.Register(moduleName, 13, "invalid validator set")
)
