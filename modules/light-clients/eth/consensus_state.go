package eth

import (
	tibccommitment "github.com/bianjieai/tibc-sdk-go/modules/core/commitment"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
)

var _ tibctypes.ConsensusState = (*ConsensusState)(nil)

func (m *ConsensusState) ClientType() string {
	return "009-eth"
}

func (m *ConsensusState) GetRoot() tibctypes.Root {
	return tibccommitment.MerkleRoot{
		Hash: m.Root,
	}
}

func (m *ConsensusState) GetTimestamp() uint64 {
	return m.Timestamp
}

func (m *ConsensusState) ValidateBasic() error {
	return nil
}
