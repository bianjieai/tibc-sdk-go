package eth

import (
	"fmt"
	"math/big"
	"time"

	tibctypes "github.com/bianjieai/tibc-sdk-go/modules/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

var (
	// Maximum number of uncles allowed in a single block
	allowedFutureBlockTimeSeconds = int64(15)
)
var _ tibctypes.Header = (*Header)(nil)

func (h Header) ClientType() string {
	return "009-eth"
}

func (h Header) GetHeight() tibctypes.Height {
	return h.Height
}

// Hash returns the block hash of the header, which is simply the keccak256 hash of its
// RLP encoding.
func (h *Header) Hash() common.Hash {
	return rlpHash(h.ToEthHeader())
}

func (h Header) ValidateBasic() error {
	// Ensure that the header's extra-data section is of a reasonable size
	if uint64(len(h.Extra)) > params.MaximumExtraDataSize {
		return fmt.Errorf("extra-data too long: %d > %d", len(h.Extra), params.MaximumExtraDataSize)
	}
	if h.Time > uint64(time.Now().Unix()+allowedFutureBlockTimeSeconds) {
		return consensus.ErrFutureBlock
	}
	// Verify that the gas limit is <= 2^63-1
	cap := uint64(0x7fffffffffffffff)
	if h.GasLimit > cap {
		return fmt.Errorf("invalid gasLimit: have %v, max %v", h.GasLimit, cap)
	}
	// Verify that the gasUsed is <= gasLimit
	if h.GasUsed > h.GasLimit {
		return fmt.Errorf("invalid gasUsed: have %d, gasLimit %d", h.GasUsed, h.GasLimit)
	}
	// Ensure that the block's difficulty is meaningful (may not be correct at this point)
	number := h.Height.RevisionHeight
	if number > 0 {
		if h.ToEthHeader().Difficulty.Uint64() == 0 {
			return tibctypes.Wrap(ErrInvalidLengthEth, "header Difficulty")
		}
	}
	return nil
}

func (h Header) ToEthHeader() EthHeader {
	difficulty := new(big.Int)
	difficulty, ok := difficulty.SetString(h.Difficulty, 10)
	if !ok {
		return EthHeader{}
	}
	baseFee := new(big.Int)
	baseFee, ok = baseFee.SetString(h.BaseFee, 10)
	if !ok {
		return EthHeader{}
	}
	return EthHeader{
		ParentHash:  common.BytesToHash(h.ParentHash),
		UncleHash:   common.BytesToHash(h.UncleHash),
		Coinbase:    common.BytesToAddress(h.Coinbase),
		Root:        common.BytesToHash(h.Root),
		TxHash:      common.BytesToHash(h.TxHash),
		ReceiptHash: common.BytesToHash(h.ReceiptHash),
		Bloom:       types.BytesToBloom(h.Bloom),
		Difficulty:  difficulty,
		Number:      big.NewInt(int64(h.Height.RevisionHeight)),
		GasLimit:    h.GasLimit,
		GasUsed:     h.GasUsed,
		Time:        h.Time,
		Extra:       h.Extra,
		MixDigest:   common.BytesToHash(h.MixDigest),
		Nonce:       types.EncodeNonce(h.Nonce),
		BaseFee:     baseFee,
	}
}
