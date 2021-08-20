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
salt: BF187E62613F8D34FF7549789E87036A
type: secp256k1

4oc/pkmJPmGDeZDWSn1LzXmpJiu+v7H7y9/UDaK+fKgtAQomv2qvmQcxilOQ/CKa
Ns0BdL8xdk/xpyvogNUzZF0XjEAv20tDfMEXMCY=
=qzos
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
salt: 37FD6B82A99A2EEFBA40FC605F2D6534
type: secp256k1

nYAzOuDvkGOGBubbcgaGqI7UocdriBsJrNT2dZVhP7sREJfpitYfKKw8I3/kQRCR
nF/1KvFwVImZB27DyB4kuqjelSqaPf5+OkOzYDY=
=6Cb0
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
salt: E4EEDC973716AD876C3A8910C1D65B20
type: secp256k1

dIdbtintSXxdAXwHT4OeZPlsrNbBgBOtwj1xq57dG9nZCk77XX7Rm20PwPy33kOj
JtDIisQLCWKU6ZtGQO5COOWDT65qaOOcpbfShzU=
=/3oB
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
	clientA := getClient(nodeURI0, grpcAddr0, chainID0, keyName0, password0, keyStore0)
	//clientB := getClient(nodeURI1, grpcAddr1, chainID1, keyName1, password1, keyStore1)
	clientC := getClient(nodeURI2, grpcAddr2, chainID2, keyName2, password2, keyStore2)
	//getjson(clientC, 4)
	//packetReceipt(clientA)
	//queryack(clientC)
	//cleanPacket(clientA, keyName0)
	updateclientTest(clientC, clientA, "testCreateClientA", keyName2)
	updateclientTest(clientA, clientC, "testCreateClientC", keyName0)
	recvCleanPacket(clientA, clientC, keyName2)

	//sendAck(clientC, clientA, keyName0)
	//bankTest(clientC)
	//queryUnreceivedPacketsAndAcks(clientA)
	//queryUnreceivedPacketsAndAcks(clientC)
	//packetReceipt(clientC)
	//updateclientTest(clientC, clientA, "testCreateClientA", keyName2)
	//packetRecive(clientA, clientC, keyName2)
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
	//res, err :=client.QueryToken("upoint")
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
	stat, err := destClient.CoreSdk.Status(context.Background())
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
	//clientState, err := client.GetClientState("testCreateClient1")
	//if err != nil {
	//	fmt.Println("GetClientState fail : ", err)
	//}
	// The trusted fields may be nil. They may be filled before relaying messages to a client.
	// The relayer is responsible for querying client and injecting appropriate trusted fields.
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
