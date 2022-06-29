package eth

import (
	commitmenttypes "github.com/bianjieai/tibc-sdk-go/modules/core/commitment"
	tibctypes "github.com/bianjieai/tibc-sdk-go/modules/types"
)

var _ tibctypes.ClientState = (*ClientState)(nil)

func (m ClientState) ClientType() string {
	return "009-eth"
}

func (m ClientState) GetLatestHeight() tibctypes.Height {
	return m.Header.Height
}

func (m ClientState) Validate() error {
	return m.Header.ValidateBasic()
}

func (m ClientState) GetDelayTime() uint64 {
	return m.TimeDelay
}

func (m ClientState) GetDelayBlock() uint64 {
	return m.BlockDelay
}

func (m ClientState) GetPrefix() tibctypes.Prefix {
	return commitmenttypes.MerklePrefix{}
}
