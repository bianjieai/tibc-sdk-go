package integration

import (
	"context"
	"fmt"
	"testing"

	tibctypes "github.com/bianjieai/tibc-sdk-go/modules/types"
	"github.com/irisnet/core-sdk-go/common/crypto"
	"github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/store"
)

const (
	nodeURIiris  = "tcp://sentry-1.mainnet.irisnet.org:26657"
	grpcAddriris = "sentry-1.mainnet.irisnet.org:9090"
	chainIDiris  = "testA"
	keyNameiris  = "chainANode0"
	passwordiris = "12345678"
	keyStoreiris = `-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: 3DD1ADDB1389FD5F06010321A9FB8230
type: sm2

1wD0rJf9qu0MMVPGqpo/fXa5VynSajIyskmhJLAkjMsWc2ssO5tSNQ0ENgTbGZbp
I6OtbT8H4FBJEH/g4aopJ+9l+zSGogNrylXlDD8=
=+MfP
-----END TENDERMINT PRIVATE KEY-----`
	chainirisLightClientName = "testCreateClientA"
	addressiris              = "iaa12z75qgsrn26cs99gvu2fq44p3ehezwdzt4durm"
)
const (
	nodeURIlocal  = "tcp://localhost:26657"
	grpcAddrlocal = "localhost:9090"
	chainIDlocal  = "testC"
	keyNamelocal  = "chainCNode0"
	passwordlocal = "12345678"
	keyStorelocal = `-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: F248AB34AE4BED4B4A5DD10E34C8BAFA
type: secp256k1

xqzXVX8JXYQWj/aJmH3hNfEqlO3IAM35C+oXkys02KfXcHZm+LDBVTRUFpONV3iB
C8FzpKaS8W69qqXhNP2sLBtHrGIk8L3T0wAVM3Q=
=2BeH
-----END TENDERMINT PRIVATE KEY-----`
	chainlocalLightClientName = "testCreateClientC"
	addresslocal              = "iaa10nfdefym9vg7c288fm4790833ee5f4p0g8w3ej"
)
const (
	nodeURI0  = "tcp://192.168.232.133:26657"
	grpcAddr0 = "192.168.232.133:9090"
	chainID0  = "testA"
	keyName0  = "chainANode0"
	password0 = "12345678"
	keyStore0 = `-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: 0CDF9AFE4EF0002010ABD3E01CBAE7E8
type: secp256k1

BNeeN1PaoRmu8jjpqA2X2EjwP6I+j02tTxU1/Z3kCijd04CV/0C28IJ3DTX/d0vN
kZkaKvjE995Gamdow95b872Sen0IyFP6FXdTbis=
=qpfu
-----END TENDERMINT PRIVATE KEY-----`
	chainALightClientName = "testCreateClientA"
	addressA              = "iaa12z75qgsrn26cs99gvu2fq44p3ehezwdzt4durm"
)

const (
	nodeURI1  = "tcp://192.168.232.135:26657"
	grpcAddr1 = "192.168.232.135:9090"
	chainID1  = "testB"
	keyName1  = "chainBNode0"
	password1 = "12345678"
	keyStore1 = `-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: 4EAEB42D9AEED740C5F9FEDA569F346C
type: secp256k1

RpPG2Ow32xTN7P3Hro34lVK8WGizKhjfyyU9QgHWFlN5Fv8zIhYeBCRsxI9QzKDX
lknIkHf36F6t06zwzYb3+ReT8cQeowsgHwOS7c8=
=YqMi
-----END TENDERMINT PRIVATE KEY-----`
	chainBLightClientName = "testCreateClientB"
	addressB              = "iaa1twsrmhpmg6rkc6l5r7rcanxeelhylxlnsw6g0l"
)
const (
	nodeURI2  = "tcp://192.168.47.128:26657"
	grpcAddr2 = "192.168.47.128:9090"
	chainID2  = "testC"
	keyName2  = "chainCNode0"
	password2 = "12345678"
	keyStore2 = `-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: B8113AE06ACB73701FBCA453D7C7AFAB
type: secp256k1

brg7lUm30ZDPDhHmSbDFWw+/pMr09rZnnUFYGiBZX2l0kG9t50Q3UkNuqxZYBdd7
MvIhB6q/n3lFY7gykW1mPi6zD6dZsgIeHK/CknY=
=6sV2
-----END TENDERMINT PRIVATE KEY-----`
	chainCLightClientName = "testCreateClientC"
	addressC              = "iaa1zvf654uuamjgecucaeklwsqeqxjguhv26f0n2g"
)

func Test_updateRinkebyEth(t *testing.T) {
	clientC, err := getIntegrationClient(nodeURI2, grpcAddr2, chainID2, keyName2, password2, keyStore2, chainCLightClientName)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error(), clientC)
		return
	}
	//getRinkebyETHjson(clientC.Tendermint)
	updateRinkebyEthClientTest(clientC, "rinkebyeth", keyName2, rinkeby)
}

func Test_getHashLen(t *testing.T) {
	header, err := GetEthNodeHeader(ethurl, 13287800)
	if err != nil {
		return
	}
	toHeader := header.ToHeader()
	fmt.Println(len(toHeader.Hash().Bytes()))
}
func Test_LocalTest(t *testing.T) {
	localClient, err := getIntegrationClient(nodeURIlocal, grpcAddrlocal, chainIDlocal, keyNamelocal, passwordlocal, keyStorelocal, chainlocalLightClientName)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error(), localClient)
		return
	}
	fmt.Println(localClient.Status(context.Background()))
}
func Test_GetIrisNetJson(t *testing.T) {
	irisClient, err := getIntegrationClient(nodeURIiris, grpcAddriris, chainIDiris, keyNameiris, passwordiris, keyStoreiris, chainALightClientName)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error(), irisClient)
		return
	}
	//getTendermintjson(irisClient.Tendermint, 11762491)
}

func Test_integrationClientTen(t *testing.T) {

	clientA, err := getIntegrationClient(nodeURI0, grpcAddr0, chainID0, keyName0, password0, keyStore0, chainALightClientName)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error(), clientA)
		return
	}
	address, err := clientA.QueryAddress(keyName0, password0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(address)
	baseTx := types.BaseTx{
		From:               keyName0,
		Gas:                0,
		Memo:               "TEST",
		Mode:               types.Commit,
		Password:           "12345678",
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}
	dec := types.NewDec(1000)
	deccoin := types.DecCoin{
		Denom:  "upoint",
		Amount: dec,
	}
	amount := types.DecCoins{deccoin}
	resu, err := clientA.Bank.Send("cosmos1mzzl97r8zkst5rgz2fyja99f3m9wh50hxg0ct9", amount, baseTx)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resu)

}

func Test_integrationClientBsc(t *testing.T) {
	clientA, err := getIntegrationClient(nodeURI0, grpcAddr0, chainID0, keyName0, password0, keyStore0, chainALightClientName)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error(), clientA)
		return
	}
	clientC, err := getIntegrationClient(nodeURI2, grpcAddr2, chainID2, keyName2, password2, keyStore2, chainCLightClientName)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error(), clientC)
		return
	}
	//getBSCjson(clientC.Tendermint)
	for i := 1; i < 500; i++ {
		fmt.Println("----------------------------------")
		updatebscclientTest(clientC, testneturl, "bsctestnet", keyName2)
	}
}

func Test_integrationClientETH(t *testing.T) {
	clientC, err := getIntegrationClient(nodeURI2, grpcAddr2, chainID2, keyName2, password2, keyStore2, chainCLightClientName)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error(), clientC)
		return
	}
	//getETHjson(clientC.Tendermint)
	//updateEthClientTest(clientC, "ethclient", keyName2, ethurl)
	for i := 1; i < 50; i++ {
		updateEthClientTest(clientC, "ethclient", keyName2, ethurl)
	}
}

func updateAllCient(clientA, clientB, clientC Client) {
	updatetendetmintclientTest(clientA, clientB, chainBLightClientName, keyName0)
	updatetendetmintclientTest(clientA, clientC, chainCLightClientName, keyName0)
	updatetendetmintclientTest(clientB, clientA, chainALightClientName, keyName1)
	updatetendetmintclientTest(clientC, clientA, chainALightClientName, keyName2)
}

func getIntegrationClient(nodeURI, grpcAddr, chainID, keyName, password, keyStore, chainName string) (Client, tibctypes.IError) {
	feeCoin, err := types.ParseDecCoins("1000upoint")
	options := []types.Option{
		types.KeyDAOOption(store.NewMemory(nil)),
		types.TimeoutOption(3000),
		types.KeyManagerOption(crypto.NewKeyManager()),
		types.BIP44PathOption(""),
		// only for irita sm2
		types.AlgoOption("sm2"),
		types.FeeOption(feeCoin),
	}
	cfg, err := types.NewClientConfig(nodeURI, grpcAddr, chainID, options...)
	if err != nil {
		return Client{}, tibctypes.New("config", 0, "error get config")
	}
	client := NewClient(cfg, chainName)
	_, err = client.Import(keyName, password, keyStore)
	if err != nil {
		return Client{}, tibctypes.New("importkey", 0, "error import key")
	}
	return client, nil
}
