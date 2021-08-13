package tendermint

import (
	"fmt"
)

const (
	SubModuleName = "tendermint-client"
	moduleName    = "tibc" + "-" + SubModuleName
)

// IBC tendermint client sentinel errors
var (
	ErrInvalidChainID         = Register(moduleName, 2, "invalid chain-id")
	ErrInvalidTrustingPeriod  = Register(moduleName, 3, "invalid trusting period")
	ErrInvalidUnbondingPeriod = Register(moduleName, 4, "invalid unbonding period")
	ErrInvalidHeaderHeight    = Register(moduleName, 5, "invalid header height")
	ErrInvalidHeader          = Register(moduleName, 6, "invalid header")
	ErrInvalidMaxClockDrift   = Register(moduleName, 7, "invalid max clock drift")
	ErrProcessedTimeNotFound  = Register(moduleName, 8, "processed time not found")
	ErrDelayPeriodNotPassed   = Register(moduleName, 9, "packet-specified delay period has not been reached")
	ErrTrustingPeriodExpired  = Register(moduleName, 10, "time since latest trusted state has passed the trusting period")
	ErrUnbondingPeriodExpired = Register(moduleName, 11, "time since latest trusted state has passed the unbonding period")
	ErrInvalidProofSpecs      = Register(moduleName, 12, "invalid proof specs")
	ErrInvalidValidatorSet    = Register(moduleName, 13, "invalid validator set")
)

type Error struct {
	codespace string
	code      uint32
	desc      string
}

func New(codespace string, code uint32, desc string) *Error {
	return &Error{codespace: codespace, code: code, desc: desc}
}

func (e Error) Error() string {
	return e.desc
}

func (e Error) ABCICode() uint32 {
	return e.code
}

func (e Error) Codespace() string {
	return e.codespace
}

func Register(codespace string, code uint32, description string) *Error {
	// TODO - uniqueness is (codespace, code) combo
	if e := getUsed(codespace, code); e != nil {
		panic(fmt.Sprintf("error with code %d is already registered: %q", code, e.desc))
	}

	err := New(codespace, code, description)
	setUsed(err)

	return err
}

var usedCodes = map[string]*Error{}

func errorID(codespace string, code uint32) string {
	return fmt.Sprintf("%s:%d", codespace, code)
}
func setUsed(err *Error) {
	usedCodes[errorID(err.codespace, err.code)] = err
}

func getUsed(codespace string, code uint32) *Error {
	return usedCodes[errorID(codespace, code)]
}
