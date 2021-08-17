package integration

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	tibc "github.com/bianjieai/tibc-sdk-go"
	tibcclient "github.com/bianjieai/tibc-sdk-go/client"
	"github.com/bianjieai/tibc-sdk-go/commitment"
	"github.com/bianjieai/tibc-sdk-go/tendermint"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	coresdk "github.com/irisnet/core-sdk-go"
	"github.com/irisnet/core-sdk-go/common/crypto"
	"github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/store"
	tenderminttypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

const (
	nodeURI0  = "tcp://192.168.232.133:26657"
	grpcAddr0 = "192.168.232.133:9090"
	chainID0  = "testA"
	keyName0  = "chainANode0"
	password0 = "12345678"
	keyStore0 = `-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: DCD05F8841AF4F0A08D59A35D7EFF238
type: secp256k1

obHSjFWeGFg4TPDxlFJHCuRHzunAD5JFKF1HYXYxHTtF1yBOHGXdLHbd2iICXyBe
/ls89K/yjvr978NEsbu3XtnVXbPvFWXincT/SGU=
=/b4+
-----END TENDERMINT PRIVATE KEY-----`
)

const (
	nodeURI1  = "tcp://192.168.232.135:26657"
	grpcAddr1 = "192.168.232.135:9090"
	chainID1  = "testB"
	keyName1  = "chainBNode0"
	password1 = "12345678"
	keyStore1 = `-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: EA4449B35A583926A711BD9929FC43B7
type: secp256k1

eStBARDqzTxKYeSLLraUzq+uB+nuw0oUkQ6qLlaqnpvylEbh5Q16tJcKto55MW6h
egjahsaUazc5y5+Ov+sFpSwACevLOU2+dryc6hU=
=7rap
-----END TENDERMINT PRIVATE KEY-----`
)
const (
	nodeURI2  = "tcp://192.168.232.140:26657"
	grpcAddr2 = "192.168.232.140:9090"
	chainID2  = "test"
	keyName2  = "node0"
	password2 = "12345678"
	keyStore2 = `-----BEGIN TENDERMINT PRIVATE KEY-----
type: secp256k1
kdf: bcrypt
salt: 24115C709F73F06EF2E88D71985C2542

JZ3Hm0AgH0eDeC0xNtJo8j8jNWVbhoeloOcQgQXxvXz5SUxOf33ssRNhPhkZ+WJC
iVfp89MmeFSpUwnOKSKWlxCLl9pygC1bEDLiPWo=
=onio
-----END TENDERMINT PRIVATE KEY-----`
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
	client := getClient(nodeURI2, grpcAddr2, chainID2, keyName2, password2, keyStore2)
	//clientB:= getClient(false)
	//	fmt.Println(client.CoreSdk.GenConn())
	getConsensusState(client, "testCreateClient", 4)
	updateclientTest(client, "testCreateClient")
	//getConsensusState(client,"testCreateClient",5)

	//	getjson(client, 4)

	//getClientState(clientA,"testCreateClientB")
	//getClientState(clientB,"testCreateClientA")

	//getconesState(client)
	//getConsensusState(client)
	//getConsensusStates(client)

}
func getClient(nodeURI, grpcAddr, chainID, keyName, password, keyStore string) tibc.Client {
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
	coreSdk := coresdk.NewClient(cfg)
	client := tibc.NewClient(coreSdk)
	_, err = client.CoreSdk.Key.Import(keyName, password, keyStore)
	if err != nil {
		panic(err)
	}
	fmt.Println(client.CoreSdk.Key.Show(keyName, "12345678"))
	return client

}

func getClientState(client tibc.Client, clientName string) {
	clientState1, err := client.GetClientState(clientName)
	if err != nil {
		panic(err)
	}
	//fmt.Println(clientState1.Validate())
	//fmt.Println(clientState1.GetLatestHeight().String())
	//fmt.Println(clientState1.ClientType())
	fmt.Println(clientState1.String())
}
func getconesState(client tibc.Client) {
	clientState1, err := client.GetClientStates()
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

func getConsensusState(client tibc.Client, clientName string, height uint64) {
	consensusState1, err := client.GetConsensusState(clientName, height)
	if err != nil {
		panic(err)
	}
	//fmt.Println(consensusState1.GetRoot().GetHash())
	//fmt.Println(consensusState1.ClientType())
	//fmt.Println(consensusState1.GetTimestamp())
	fmt.Println(consensusState1.String())

}
func getConsensusStates(client tibc.Client) {
	consensusState1, err := client.GetConsensusStates("testCreateClient")
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

func bankTest(client tibc.Client) {
	coins, _ := types.ParseDecCoins("100stake")
	to := "iaa12a08j6scetjx8kesf6t0guh2jhe0a5c5zhae2m"
	baseTx := types.BaseTx{
		From:               keyName0,
		Gas:                0,
		Memo:               "TEST",
		Mode:               types.Commit,
		Password:           "12345678",
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}
	res, err := client.CoreSdk.Bank.Send(to, coins, baseTx)
	//res, err :=client.QueryToken("upoint")
	if err != nil {
		fmt.Println("Bank Send fail : ", err)
	}
	fmt.Println(res)
}

//destClient tibc.Client,
func updateclientTest(sourceClient tibc.Client, chainName string) {
	baseTx := types.BaseTx{
		From:               keyName2,
		Gas:                1000000,
		Memo:               "TEST",
		Mode:               types.Commit,
		Password:           "12345678",
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}
	request := tibctypes.UpdateClientRequest{
		ChainName: chainName,
		Header:    CreateHeader(sourceClient, 5),
	}

	ress, err := sourceClient.UpdateClient(request, baseTx)
	if err != nil {
		fmt.Println("UpdateClient fail :", err)
	}
	fmt.Println(ress)
}

func CreateHeader(client tibc.Client, height int64) *tendermint.Header {

	res, err := client.CoreSdk.QueryBlock(height)
	if err != nil {
		fmt.Println("QueryBlock fail:  ", err)
	}
	tmHeader := res.Block.Header

	trustHeight := tibcclient.NewHeight(0, 4)
	rescommit, err := client.CoreSdk.Commit(context.Background(), &res.BlockResult.Height)
	commit := rescommit.Commit
	signedHeader := &tenderminttypes.SignedHeader{
		Header: tmHeader.ToProto(),
		Commit: commit.ToProto(),
	}
	clientState, err := client.GetClientState("testCreateClient")
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

func queryValidatorSet(height int64, client tibc.Client) *tenderminttypes.ValidatorSet {

	validators, err := client.CoreSdk.Validators(context.Background(), &height, nil, nil)
	if err != nil {
		fmt.Println("queryValidatorSet fail :", err)
	}
	validatorSet, err := tmtypes.NewValidatorSet(validators.Validators).ToProto()
	if err != nil {
		fmt.Println("queryValidatorSet fail :", err)
	}
	return validatorSet
}

func queryValidatorSet1(height int64, client tibc.Client) *tmtypes.ValidatorSet {
	validators, err := client.CoreSdk.Validators(context.Background(), &height, nil, nil)
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
func getjson(client tibc.Client, height int64) {

	//ClientState
	var fra = tendermint.Fraction{
		Numerator:   1,
		Denominator: 3,
	}
	res, err := client.CoreSdk.QueryBlock(height)
	if err != nil {
		fmt.Println("QueryBlock fail:  ", err)
	}
	tmHeader := res.Block.Header
	fmt.Println(tmHeader.ChainID)
	lastHeight := tibcclient.NewHeight(0, 4)
	var clientstate = &tendermint.ClientState{
		ChainId:         tmHeader.ChainID,
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
	var consensusState = &tendermint.ConsensusState{
		Timestamp:          time.Date(2021, 8, 6, 0, 0, 0, 0, time.UTC),
		Root:               commitment.NewMerkleRoot([]byte("app_hash")),
		NextValidatorsHash: queryValidatorSet1(res.Block.Height, client).Hash(),
	}

	b0, err := client.Marshaler.MarshalJSON(clientstate)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("client_state.json", b0, os.ModeAppend)
	if err != nil {
		return
	}
	b1, err := client.Marshaler.MarshalJSON(consensusState)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("consensus_state.json", b1, os.ModeAppend)
	if err != nil {
		return
	}

}
