package eth

import (
	commitmenttypes "github.com/bianjieai/tibc-sdk-go/commitment"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
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
	//todo ? add delaytime
	return 0
}

func (m ClientState) GetDelayBlock() uint64 {
	//todo ? add delayblock
	return 0
}

func (m ClientState) GetPrefix() tibctypes.Prefix {
	return commitmenttypes.MerklePrefix{}
}
