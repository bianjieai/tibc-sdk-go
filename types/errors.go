package types

import "fmt"

const RootCodeSpace = "tibc-sdk-go"

var (
	ErrInternal                     = Register(RootCodeSpace, 1, "internal")
	ErrChainConn                    = Register(RootCodeSpace, 2, "connection chain failed")
	ErrGetLightClientState          = Register(RootCodeSpace, 3, "failed to get light client state")
	ErrGetBlockHeader               = Register(RootCodeSpace, 4, "failed to get block header")
	ErrUpdateClient                 = Register(RootCodeSpace, 5, "failed to update client")
	ErrGetPackets                   = Register(RootCodeSpace, 6, "failed to get packets")
	ErrGetCommitmentPacket          = Register(RootCodeSpace, 7, "failed to get commitment packet")
	ErrGetAckPacket                 = Register(RootCodeSpace, 8, "failed to get ack packet")
	ErrGetReceiptPacket             = Register(RootCodeSpace, 9, "failed to get receipt packet")
	ErrGetProof                     = Register(RootCodeSpace, 10, "failed to get proof")
	ErrGetLatestHeight              = Register(RootCodeSpace, 11, "failed to get latest height")
	ErrRecvPacket                   = Register(RootCodeSpace, 12, "failed to recv packet")
	ErrNotProduced                  = Register(RootCodeSpace, 13, "failed to not produced")
	ErrUnknownMsg                   = Register(RootCodeSpace, 14, "failed to unknown msg type")
	ErrUnpackAny                    = Register(RootCodeSpace, 15, "failed to unpack any")
	ErrGetLightClientConsensusState = Register(RootCodeSpace, 16, "failed to light consensus state")
	ErrGetRelayer                   = Register(RootCodeSpace, 17, "failed to get relayer")
	ErrGetAddress                   = Register(RootCodeSpace, 18, "failed to get address")
	ErrPackAny                      = Register(RootCodeSpace, 19, "failed to pack any")
	ErrGetUnreceivedPacket          = Register(RootCodeSpace, 20, "failed to get unreceived packet")
	ErrSendAckPacket                = Register(RootCodeSpace, 21, "failed to send ack packet")
	ErrSendCleanPacket              = Register(RootCodeSpace, 22, "failed to send clean packet")
	ErrRecvCleanPacket              = Register(RootCodeSpace, 23, "failed to recv clean packet")
)

var usedCodes = map[string]*Error{}

func getUsed(codespace string, code uint32) *Error {
	return usedCodes[errorID(codespace, code)]
}

func setUsed(err *Error) {
	usedCodes[errorID(err.codeSpace, err.code)] = err
}

func errorID(codespace string, code uint32) string {
	return fmt.Sprintf("%s:%d", codespace, code)
}

type IError interface {
	error
	Code() uint32
	Codespace() string
}

type Error struct {
	codeSpace string
	code      uint32
	desc      string
}

func (e Error) Codespace() string {
	return e.codeSpace
}

func New(codeSpace string, code uint32, desc string) *Error {
	return &Error{codeSpace: codeSpace, code: code, desc: desc}
}

func (e Error) Error() string {
	return e.desc
}

func (e Error) Code() uint32 {
	return e.code
}

func Register(codespace string, code uint32, description string) *Error {
	if e := getUsed(codespace, code); e != nil {
		panic(fmt.Sprintf("error with code %d is already registered: %q", code, e.desc))
	}

	err := New(codespace, code, description)
	setUsed(err)

	return err
}

type WrappedError struct {
	// This error layer description.
	msg string
	// The underlying error that triggered this one.
	parent error
}

func (e *WrappedError) Error() string {
	return fmt.Sprintf("%s: %s", e.msg, e.parent.Error())
}

func (e *WrappedError) Cause() error {
	return e.parent
}

// Is reports whether any error in e's chain matches a target.
func (e *WrappedError) Is(target error) bool {
	if e == target {
		return true
	}

	w := e.Cause()
	for {
		if w == target {
			return true
		}

		x, ok := w.(causer)
		if ok {
			w = x.Cause()
		}
		if x == nil {
			return false
		}
	}
}

// Unwrap implements the built-in errors.Unwrap
func (e *WrappedError) Unwrap() error {
	return e.parent
}

// causer is an interface implemented by an error that supports wrapping. Use
// it to test if an error wraps another error instance.
type causer interface {
	Cause() error
}

type unpacker interface {
	Unpack() []error
}
