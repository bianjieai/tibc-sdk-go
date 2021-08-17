package tendermint

import (
	"fmt"

	commitmenttypes "github.com/bianjieai/tibc-sdk-go/commitment"
	coresdk "github.com/irisnet/core-sdk-go"
	"github.com/irisnet/core-sdk-go/common/codec"

)

// QueryTendermintProof performs an ABCI query with the given key and returns
// the value of the query, the proto encoded merkle proof, and the height of
// the Tendermint block containing the state root. The desired tendermint height
// to perform the query should be set in the client context. The query will be
// performed at one below this height (at the IAVL version) in order to obtain
// the correct merkle proof. Proof queries at height less than or equal to 2 are
// not supported. Queries with a client context height of 0 will perform a query
// at the lastest state available.
// Issue: https://github.com/cosmos/cosmos-sdk/issues/6567
func QueryTendermintProof(coreClient coresdk.Client, height int64, key []byte) ([]byte, []byte, uint64, error) {

	// ABCI queries at heights 1, 2 or less than or equal to 0 are not supported.
	// Base app does not support queries for height less than or equal to 1.
	// Therefore, a query at height 2 would be equivalent to a query at height 3.
	// A height of 0 will query with the lastest state.
	if height != 0 && height <= 2 {
		return nil, nil,0, fmt.Errorf("proof queries at height <= 2 are not supported")
	}
	// Use the IAVL height if a valid tendermint height is passed in.
	// A height of 0 will query with the latest state.
	if height != 0 {
		height--
	}
	res, err := coreClient.QueryStore(key,"tibc",height,true)
	if err != nil {
		return nil, nil, 0, err
	}

	merkleProof, err := commitmenttypes.ConvertProofs(res.ProofOps)
	if err != nil {
		return nil, nil, 0, err
	}
	cdc := codec.NewProtoCodec(coreClient.EncodingConfig().InterfaceRegistry)

	proofBz, err := cdc.MarshalBinaryBare(&merkleProof)
	if err != nil {
		return nil, nil, 0, err
	}

	return res.Value, proofBz,  uint64(res.Height)+1, nil
}
