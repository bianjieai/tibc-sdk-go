package bsc

import (
	"math/big"

	tibctypes "github.com/bianjieai/tibc-sdk-go/modules/types"
	"github.com/ethereum/go-ethereum/common"
)

var _ tibctypes.Header = (*Header)(nil)

func (h Header) ClientType() string {
	return "008-bsc"
}

func (h Header) GetHeight() tibctypes.Height {
	return h.Height
}

// Hash returns the block hash of the header, which is simply the keccak256 hash of its
// RLP encoding.
func (h *Header) Hash() common.Hash {
	return rlpHash(h.ToBscHeader())
}

func (h Header) ValidateBasic() error {
	number := h.Height.RevisionHeight

	// Check that the extra-data contains the vanity, validators and signature.
	if len(h.Extra) < extraVanity {
		return tibctypes.Wrap(ErrMissingVanity, "header Extra")
	}
	if len(h.Extra) < extraVanity+extraSeal {
		return tibctypes.Wrap(ErrMissingSignature, "header Extra")
	}

	// Ensure that the mix digest is zero as we don't have fork protection currently
	if common.BytesToHash(h.MixDigest) != (common.Hash{}) {
		return tibctypes.Wrap(ErrInvalidMixDigest, "header MixDigest")
	}
	// Ensure that the block doesn't contain any uncles which are meaningless in PoA
	if common.BytesToHash(h.UncleHash) != uncleHash {
		return tibctypes.Wrap(ErrInvalidUncleHash, "header UncleHash")
	}
	// Ensure that the block's difficulty is meaningful (may not be correct at this point)
	if number > 0 {
		if h.Difficulty == 0 {
			return tibctypes.Wrap(ErrInvalidDifficulty, "header Difficulty")
		}
	}
	return nil
}

func (h Header) ToBscHeader() BscHeader {
	return BscHeader{
		ParentHash:  common.BytesToHash(h.ParentHash),
		UncleHash:   common.BytesToHash(h.UncleHash),
		Coinbase:    common.BytesToAddress(h.Coinbase),
		Root:        common.BytesToHash(h.Root),
		TxHash:      common.BytesToHash(h.TxHash),
		ReceiptHash: common.BytesToHash(h.ReceiptHash),
		Bloom:       BytesToBloom(h.Bloom),
		Difficulty:  big.NewInt(int64(h.Difficulty)),
		Number:      big.NewInt(int64(h.Height.RevisionHeight)),
		GasLimit:    h.GasLimit,
		GasUsed:     h.GasUsed,
		Time:        h.Time,
		Extra:       h.Extra,
		MixDigest:   common.BytesToHash(h.MixDigest),
		Nonce:       BytesToBlockNonce(h.Nonce),
	}
}
