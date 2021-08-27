package integration

import (
	"context"
	"fmt"
	"testing"

	tibc "github.com/bianjieai/tibc-sdk-go"
	tibcclient "github.com/bianjieai/tibc-sdk-go/client"
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
salt: D145AFB29D332974833DB1AD5912D3D1
type: secp256k1

KCtCKbIq/EVtNGdZSoa4cwkLfMk9Z7tMQ6P1DrSe0KqkmMutzA2hh3WbtrSdJsgD
KZM9JwHI4ROYyXGsTD5s8oL+kjErG7zcpjhbX7g=
=K8L7
-----END TENDERMINT PRIVATE KEY-----`
)

const (
	nodeURI1  = "tcp://192.168.232.135:26657"
	grpcAddr1 = "192.168.232.135:9090"
	chainID1  = "testB"
	keyName1  = "chainBNode0"
	password1 = "12345678"
	keyStore1 = `-----BEGIN TENDERMINT PRIVATE KEY-----
salt: 2F4473E8BA06E6142C476C009FF72423
type: secp256k1
kdf: bcrypt

T+2XCYRqkZBSFXGTToW5ryGFizmWGirJMPy6Dcc7KB6GOM6aJrASCa8s4Zi34zDp
EYxOsdB0TPJtfpFWRDaMZIxwiyXowgCLD09FOps=
=OZa+
-----END TENDERMINT PRIVATE KEY-----`
)
const (
	nodeURI2  = "tcp://192.168.232.140:26657"
	grpcAddr2 = "192.168.232.140:9090"
	chainID2  = "testC"
	keyName2  = "chainCNode0"
	password2 = "12345678"
	keyStore2 = `-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: 25CBC5EF99C0D6831E61F084317B6CAB
type: secp256k1

VJU1N6MjKQTeOnuQnWEv3He5Ygv1JahIUvDO1pTK7QB7199+ftNMS2eOTV7iNJGb
Un9YJEuYJHpP/1cz2VuOuMEgqXFYe/Z6uWfGuMs=
=UJ7v
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
	clientA, err := getClient(nodeURI0, grpcAddr0, chainID0, keyName0, password0, keyStore0)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error())
		return
	}
	clientB, err := getClient(nodeURI1, grpcAddr1, chainID1, keyName1, password1, keyStore1)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error())
		return
	}

	updateclientTest(clientA, clientB, "testCreateClientB", "testCreateClientB")

}

func getClient(nodeURI, grpcAddr, chainID, keyName, password, keyStore string) (tibc.Client, *tibctypes.Error) {
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
		return tibc.Client{}, tibctypes.New("config", 0, "error get config")
	}
	coreSdk := coresdk.NewClient(cfg)
	client := tibc.NewClient(coreSdk)
	_, err = client.CoreSdk.Key.Import(keyName, password, keyStore)
	if err != nil {

		return tibc.Client{}, tibctypes.New("importkey", 0, "error import key")
	}
	return client, nil

}

func getClientState(client tibc.Client, clientName string) {
	clientState1, err := client.GetClientState(clientName)
	if err != nil {
		panic(err)
	}
	fmt.Println(clientState1.String())
}
func getClientStates(client tibc.Client) {
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
	fmt.Println(consensusState1.String())

}
func getConsensusStates(client tibc.Client) {
	consensusState1, err := client.GetConsensusStates("testCreateClient")
	if err != nil {
		panic(err)
	}
	fmt.Println("consensusState: ")
	for _, value := range consensusState1 {
		if value == nil {
			break
		}
		fmt.Println(value.String())
	}
}

func bankTest(client tibc.Client) {
	coins, _ := types.ParseDecCoins("100stake")
	to := "iaa1mlj9nsud3d9yaccgymf4ay9yckr268qttggnrj"
	baseTx := types.BaseTx{
		From:               keyName2,
		Gas:                0,
		Memo:               "TEST",
		Mode:               types.Commit,
		Password:           "12345678",
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}
	res, err := client.CoreSdk.Bank.Send(to, coins, baseTx)
	if err != nil {
		fmt.Println("Bank Send fail : ", err)
	}
	fmt.Println(res)
}

//destClient tibc.Client,
func updateclientTest(sourceClient tibc.Client, destClient tibc.Client, chainName, keyname string) {
	baseTx := types.BaseTx{
		From:               keyname,
		Gas:                0,
		Memo:               "TEST",
		Mode:               types.Commit,
		Password:           "12345678",
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}
	lightClientState, err := sourceClient.GetClientState(chainName)
	if err != nil {
		fmt.Println("UpdateClient fail :", err)
		return
	}
	height := int64(lightClientState.GetLatestHeight().GetRevisionHeight())
	stat, err1 := destClient.CoreSdk.Status(context.Background())
	if err1 != nil {
		fmt.Println("get Status fail :", err1)
		return
	}
	height1 := stat.SyncInfo.LatestBlockHeight
	request := tibctypes.UpdateClientRequest{
		ChainName: chainName,
		Header:    CreateHeader(destClient, height1, tibcclient.NewHeight(0, uint64(height)), lightClientState),
	}

	ress, err := sourceClient.UpdateClient(request, baseTx)
	if err != nil {
		fmt.Println("UpdateClient fail :", err)
		return
	}
	fmt.Println(ress)
}

func CreateHeader(client tibc.Client, height int64, trustHeight tibcclient.Height, clientState tibctypes.ClientState) *tendermint.Header {

	res, err := client.CoreSdk.QueryBlock(height)
	if err != nil {
		fmt.Println("QueryBlock fail:  ", err)
	}
	tmHeader := res.Block.Header

	rescommit, err := client.CoreSdk.Commit(context.Background(), &res.BlockResult.Height)
	commit := rescommit.Commit
	signedHeader := &tenderminttypes.SignedHeader{
		Header: tmHeader.ToProto(),
		Commit: commit.ToProto(),
	}
	return &tendermint.Header{
		SignedHeader:      signedHeader,
		ValidatorSet:      queryValidatorSet(height, client),
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
