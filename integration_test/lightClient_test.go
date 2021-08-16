package integration

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	tibcclient "github.com/bianjieai/tibc-sdk-go/client"
	"github.com/bianjieai/tibc-sdk-go/commitment"
	"github.com/bianjieai/tibc-sdk-go/tendermint"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	"github.com/irisnet/core-sdk-go/common/crypto"
	"github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/store"
	tenderminttypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

const (
	nodeURI  = "tcp://192.168.232.140:26657"
	grpcAddr = "192.168.232.140:9090"
	chainID  = "test"
	keyName  = "node0"
	password = "12345678"
	keyStore = `-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: C9C670FEAF9DDD4941B95B4FB80A4A0D
type: secp256k1

WkFdsL0l55MTvvfh9JUk1mssQR5LOiiT9B56A3pksQXJK5OZm7cW9bI7l7zAZYij
URKMjHoEjcZRyaYWIVNsfx3EJGVK+zHcONNESn4=
=eJ3c
-----END TENDERMINT PRIVATE KEY-----
`
)

type TokenManager struct{}

func (TokenManager TokenManager) QueryToken(denom string) (types.Token, error) {
	return types.Token{}, nil
}

func (TokenManager TokenManager) SaveTokens(tokens ...types.Token) {
	return
}

func (TokenManager TokenManager) ToMinCoin(coins ...types.DecCoin) (types.Coins, types.Error) {
	for i := range coins {
		if coins[i].Denom == "iris" {
			coins[i].Denom = "uiris"
			coins[i].Amount = coins[i].Amount.MulInt(types.NewIntWithDecimal(1, 6))
		}
	}
	ucoins, _ := types.DecCoins(coins).TruncateDecimal()
	return ucoins, nil
}

func (TokenManager TokenManager) ToMainCoin(coins ...types.Coin) (types.DecCoins, types.Error) {
	decCoins := make(types.DecCoins, len(coins), 0)
	for _, coin := range coins {
		if coin.Denom == "uiris" {
			amtount := types.NewDecFromInt(coin.Amount).Mul(types.NewDecWithPrec(1, 6))
			decCoins = append(decCoins, types.NewDecCoinFromDec("iris", amtount))
		}
	}
	return decCoins, nil
}

func Test_ClientCreat(t *testing.T) {

	feeCoin, err := types.ParseDecCoins("10stake")
	options := []types.Option{
		types.KeyDAOOption(store.NewMemory(nil)),
		types.TimeoutOption(10),
		types.TokenManagerOption(TokenManager{}),
		types.KeyManagerOption(crypto.NewKeyManager()),
		types.BIP44PathOption(""),
		types.FeeOption(feeCoin),
	}

	cfg, err := types.NewClientConfig(nodeURI, grpcAddr, chainID, options...)
	if err != nil {
		panic(err)
	}
	client := newClient(cfg)
	_, err = client.Key.Import(keyName, password, keyStore)
	if err != nil {
		panic(err)
	}
	fmt.Println(client.Key.Show("node0", "12345678"))
	getclientstate(client, 66)
}

func getClientState(client clientforlightclient) {
	clientState1, err := client.TendermintClient.GetClientState("testCreateClient1")
	if err != nil {
		panic(err)
	}
	fmt.Println(clientState1.Validate())
	fmt.Println(clientState1.GetLatestHeight().String())
	fmt.Println(clientState1.ClientType())
	fmt.Println(clientState1.String())
}
func getconesState(client clientforlightclient) {
	clientState1, err := client.TendermintClient.GetClientStates()
	if err != nil {
		panic(err)
	}
	for _, value := range clientState1 {
		if value == nil {
			break
		}
		fmt.Println(value.String())
	}
}

func getConsensusState(client clientforlightclient) {
	consensusState1, err := client.TendermintClient.GetConsensusState("testCreateClient1", 8)
	if err != nil {
		panic(err)
	}
	fmt.Println(consensusState1.GetRoot().GetHash())
	fmt.Println(consensusState1.ClientType())
	fmt.Println(consensusState1.GetTimestamp())
	fmt.Println(consensusState1.String())

}
func getConsensusStates(client clientforlightclient) {
	consensusState1, err := client.TendermintClient.GetConsensusStates("testCreateClient1")
	if err != nil {
		panic(err)
	}
	for _, value := range consensusState1 {
		if value == nil {
			break
		}
		fmt.Println(value.String())
	}
}

func updateclientTest(client clientforlightclient) {
	baseTx := types.BaseTx{
		From:               "node0",
		Gas:                0,
		Memo:               "TEST",
		Mode:               types.Commit,
		Password:           "12345678",
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}

	request := tibctypes.UpdateClientRequest{
		ChainName: "testCreateClient1",
		Header:    CreateHeader(client, 66),
	}

	ress, err := client.TendermintClient.UpdateClient(request, baseTx)
	//res, err :=client.QueryToken("upoint")
	if err != nil {
		fmt.Println("UpdateClient fail :", err)
	}
	fmt.Println(ress)
}

func bankTest(client clientforlightclient) {
	coins, _ := types.ParseDecCoins("100stake")
	to := "iaa13uk9hnts6ajm5xdcvr9k97v24wu0tgu2wmkwsm"
	baseTx := types.BaseTx{
		From:               "node0",
		Gas:                0,
		Memo:               "TEST",
		Mode:               types.Commit,
		Password:           "12345678",
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}
	res, err := client.Bank.Send(to, coins, baseTx)
	//res, err :=client.QueryToken("upoint")
	if err != nil {
		fmt.Println("Bank Send fail : ", err)
	}
	fmt.Println(res)
}
func CreateHeader(client clientforlightclient, height int64) *tendermint.Header {

	res, err := client.QueryBlock(height)
	if err != nil {
		fmt.Println("QueryBlock fail:  ", err)
	}

	tmHeader := tmtypes.Header{
		Version:            res.Block.Version,
		ChainID:            chainID,
		Height:             res.Block.Height,
		Time:               res.Block.Time,
		LastBlockID:        res.Block.LastBlockID,
		LastCommitHash:     res.Block.LastCommitHash,
		DataHash:           res.Block.DataHash,
		ValidatorsHash:     res.Block.ValidatorsHash,
		NextValidatorsHash: res.Block.NextValidatorsHash,
		ConsensusHash:      res.Block.ConsensusHash,
		AppHash:            res.Block.AppHash,
		LastResultsHash:    res.Block.LastResultsHash,
		EvidenceHash:       res.Block.EvidenceHash,
		ProposerAddress:    res.Block.ProposerAddress, //nolint:staticcheck
	}
	trustHeight := tibcclient.NewHeight(0, 4)
	rescommit, err := client.Commit(context.Background(), &res.BlockResult.Height)
	commit := rescommit.Commit
	signedHeader := &tenderminttypes.SignedHeader{
		Header: tmHeader.ToProto(),
		Commit: commit.ToProto(),
	}
	clientState, err := client.TendermintClient.GetClientState("testCreateClient1")
	if err != nil {
		fmt.Println("GetClientState fail : ", err)
	}
	// The trusted fields may be nil. They may be filled before relaying messages to a client.
	// The relayer is responsible for querying client and injecting appropriate trusted fields.
	return &tendermint.Header{
		SignedHeader:      signedHeader,
		ValidatorSet:      queryValidatorSet(res.Block.Height, client),
		TrustedHeight:     trustHeight,
		TrustedValidators: queryValidatorSet(int64(clientState.GetLatestHeight().GetRevisionHeight()), client),
	}
}

func queryValidatorSet(height int64, client clientforlightclient) *tenderminttypes.ValidatorSet {

	validators, err := client.Validators(context.Background(), &height, nil, nil)
	if err != nil {
		fmt.Println("queryValidatorSet fail :", err)
	}
	validatorSet, err := tmtypes.NewValidatorSet(validators.Validators).ToProto()
	if err != nil {
		fmt.Println("queryValidatorSet fail :", err)
	}
	return validatorSet
}

func queryValidatorSet1(height int64, client clientforlightclient) *tmtypes.ValidatorSet {
	validators, err := client.Validators(context.Background(), &height, nil, nil)
	if err != nil {
		fmt.Println("queryValidatorSet1 fail :", err)
	}
	validatorSet := tmtypes.NewValidatorSet(validators.Validators)

	return validatorSet
}

//Generate a JSON file needed to create the light client
//Add the following string to the header after the file is generated
//"@type":"/tibc.lightclients.tendermint.v1.ClientState",
//"@type":"/tibc.lightclients.tendermint.v1.ConsensusState",
func getclientstate(client clientforlightclient, height int64) {

	//ClientState
	var fra = tendermint.Fraction{
		Numerator:   1,
		Denominator: 3,
	}
	lastHeight := tibcclient.NewHeight(0, 4)

	var clientstate = &tendermint.ClientState{
		ChainId:         chainID,
		TrustLevel:      fra,
		TrustingPeriod:  time.Hour * 24 * 7 * 2,
		UnbondingPeriod: time.Hour * 24 * 7 * 3,
		MaxClockDrift:   time.Second * 10,
		LatestHeight:    lastHeight,
		ProofSpecs:      commitment.GetSDKSpecs(),
		MerklePrefix:    commitment.MerklePrefix{KeyPrefix: []byte("ibc")},
		TimeDelay:       0,
	}
	//ConsensusState
	res, err := client.QueryBlock(height)
	var consensusState = &tendermint.ConsensusState{
		Timestamp:          time.Date(2021, 8, 6, 0, 0, 0, 0, time.UTC),
		Root:               commitment.NewMerkleRoot([]byte("app_hash")),
		NextValidatorsHash: queryValidatorSet1(res.Block.Height, client).Hash(),
	}

	b0, err := client.TendermintClient.Marshaler.MarshalJSON(clientstate)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("ClientState.json", b0, os.ModeAppend)
	if err != nil {
		return
	}
	b1, err := client.TendermintClient.Marshaler.MarshalJSON(consensusState)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("ConsensusState.json", b1, os.ModeAppend)
	if err != nil {
		return
	}

}
