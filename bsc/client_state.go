package bsc

import (
	commitmenttypes "github.com/bianjieai/tibc-sdk-go/commitment"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
)

var _ tibctypes.ClientState = (*ClientState)(nil)

func (m ClientState) ClientType() string {
	return "008-bsc"
}

func (m ClientState) GetLatestHeight() tibctypes.Height {
	return m.Header.Height
}

func (m ClientState) Validate() error {
	return m.Header.ValidateBasic()
}

func (m ClientState) GetDelayTime() uint64 {
	return uint64((2*len(m.Validators)/3 + 1)) * m.BlockInteval
}

func (m ClientState) GetDelayBlock() uint64 {
	return uint64(2*len(m.Validators)/3 + 1)
}

func (m ClientState) GetPrefix() tibctypes.Prefix {
	return commitmenttypes.MerklePrefix{}
}
