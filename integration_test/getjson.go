package integration

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	tibc "github.com/bianjieai/tibc-sdk-go"
	tibcclient "github.com/bianjieai/tibc-sdk-go/client"
	"github.com/bianjieai/tibc-sdk-go/commitment"
	"github.com/bianjieai/tibc-sdk-go/tendermint"
	tmtypes "github.com/tendermint/tendermint/types"
)

const (
	ConType = "{\"@type\":\"/tibc.lightclients.tendermint.v1.ConsensusState\","
	StaType = "{\"@type\":\"/tibc.lightclients.tendermint.v1.ClientState\","
)

//Generate a JSON file needed to create the light client
//Add the following string to the header after the file is generated
//"@type":"/tibc.lightclients.tendermint.v1.ClientState",
//"@type":"/tibc.lightclients.tendermint.v1.ConsensusState",
func getjson(client tibc.Client, height int64) {

	//ClientState
	var fra = tendermint.Fraction{
		Numerator:   1,
		Denominator: 3,
	}
	res, err := client.QueryBlock(height)
	if err != nil {
		fmt.Println("QueryBlock fail:  ", err)
	}
	tmHeader := res.Block.Header
	lastHeight := tibcclient.NewHeight(0, 4)
	var clientstate = &tendermint.ClientState{
		ChainId:         tmHeader.ChainID,
		TrustLevel:      fra,
		TrustingPeriod:  time.Hour * 24 * 7 * 2,
		UnbondingPeriod: time.Hour * 24 * 7 * 3,
		MaxClockDrift:   time.Second * 10,
		LatestHeight:    lastHeight,
		ProofSpecs:      commitment.GetSDKSpecs(),
		MerklePrefix:    commitment.MerklePrefix{KeyPrefix: []byte("tibc")},
		TimeDelay:       0,
	}
	//ConsensusState
	var consensusState = &tendermint.ConsensusState{
		Timestamp:          tmHeader.Time,
		Root:               commitment.NewMerkleRoot([]byte("app_hash")),
		NextValidatorsHash: queryValidatorSet1(res.Block.Height, client).Hash(),
	}

	b0, err := client.Marshaler.MarshalJSON(clientstate)
	if err != nil {
		panic(err)
	}
	b0 = []byte(StaType + string(b0)[1:])
	clientStateName := tmHeader.ChainID + "_client_state.json"
	err = ioutil.WriteFile(clientStateName, b0, os.ModeAppend)
	if err != nil {
		return
	}
	b1, err := client.Marshaler.MarshalJSON(consensusState)
	if err != nil {
		panic(err)
	}
	b1 = []byte(ConType + string(b1)[1:])
	clientConsensusStateName := tmHeader.ChainID + "_consensus_state.json"
	err = ioutil.WriteFile(clientConsensusStateName, b1, os.ModeAppend)
	if err != nil {
		return
	}
}

func queryValidatorSet1(height int64, client tibc.Client) *tmtypes.ValidatorSet {
	validators, err := client.Validators(context.Background(), &height, nil, nil)
	if err != nil {
		fmt.Println("queryValidatorSet1 fail :", err)
	}
	validatorSet := tmtypes.NewValidatorSet(validators.Validators)
	if err != nil {
		fmt.Println("queryValidatorSet fail :", err)
	}
	return validatorSet
}
