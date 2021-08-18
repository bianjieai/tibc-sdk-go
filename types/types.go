package types

import (
	"errors"

	coretypes "github.com/irisnet/core-sdk-go/common/codec/types"
)


// UnpackHeader unpacks an Any into a Header. It returns an error if the
// consensus state can't be unpacked into a Header.
func UnpackHeader(any *coretypes.Any) (Header, error) {
	if any == nil {
		return nil, errors.New("protobuf Any message cannot be nil")
	}

	header, ok := any.GetCachedValue().(Header)
	if !ok {
		return nil, errors.New("cannot unpack Any into Header")
	}
	return header, nil
}
