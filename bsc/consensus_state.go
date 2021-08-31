package bsc

import (
	commitmenttypes "github.com/bianjieai/tibc-sdk-go/commitment"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
)

var _ tibctypes.ConsensusState = (*ConsensusState)(nil)

func (m *ConsensusState) ClientType() string {
	return "008-bsc"
}

func (m *ConsensusState) GetRoot() tibctypes.Root {
	return commitmenttypes.MerkleRoot{
		Hash: m.Root,
	}
}

func (m *ConsensusState) GetTimestamp() uint64 {
	return m.Timestamp
}

func (m *ConsensusState) ValidateBasic() error {
	return nil
}
