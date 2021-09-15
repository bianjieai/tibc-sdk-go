package integration

import (
	"fmt"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	"github.com/irisnet/core-sdk-go/common/crypto"
	"github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/store"
	"testing"
)

const (
	nodeURIiris  = "tcp://sentry-1.mainnet.irisnet.org:26657"
	grpcAddriris = "sentry-1.mainnet.irisnet.org:9090"
	chainIDiris  = "testA"
	keyNameiris  = "chainANode0"
	passwordiris = "12345678"
	keyStoreiris = `-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: 0CDF9AFE4EF0002010ABD3E01CBAE7E8
type: secp256k1

BNeeN1PaoRmu8jjpqA2X2EjwP6I+j02tTxU1/Z3kCijd04CV/0C28IJ3DTX/d0vN
kZkaKvjE995Gamdow95b872Sen0IyFP6FXdTbis=
=qpfu
-----END TENDERMINT PRIVATE KEY-----`
	chainirisLightClientName = "testCreateClientA"
	addressiris              = "iaa12z75qgsrn26cs99gvu2fq44p3ehezwdzt4durm"
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
salt: AA6711769B1862B00F24BE3A52E21AA8
type: secp256k1

h2N+FyIkP4ajJxoPHMi8Sm/bsTb7bkkpjKq3M+/m9EbeHh297ezgtbR0S1UP40y0
Rkw5YDP21pxbcgC2yCgyEOncHly1tIui8fkJKho=
=L9VU
-----END TENDERMINT PRIVATE KEY-----`
	chainCLightClientName = "testCreateClientC"
	addressC              = "iaa10nfdefym9vg7c288fm4790833ee5f4p0g8w3ej"
)

func Test_GetIrisNetJson(t *testing.T) {
	irisClient, err := getIntegrationClient(nodeURIiris, grpcAddriris, chainIDiris, keyNameiris, passwordiris, keyStoreiris, chainALightClientName)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error(), irisClient)
		return
	}
	getTendermintjson(irisClient.Tendermint, 11762491)
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
		Denom:  "stake",
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
	updateEthClientTest(clientC, "ethclient", keyName2)
	//for i := 1; i < 500; i++ {
	//	fmt.Println("----------------------------------")
	//	updateEthClientTest(clientC, "ethclient", keyName2)
	//}
}

func updateAllCient(clientA, clientB, clientC Client) {
	updatetendetmintclientTest(clientA, clientB, chainBLightClientName, keyName0)
	updatetendetmintclientTest(clientA, clientC, chainCLightClientName, keyName0)
	updatetendetmintclientTest(clientB, clientA, chainALightClientName, keyName1)
	updatetendetmintclientTest(clientC, clientA, chainALightClientName, keyName2)
}

func getIntegrationClient(nodeURI, grpcAddr, chainID, keyName, password, keyStore, chainName string) (Client, tibctypes.IError) {
	feeCoin, err := types.ParseDecCoins("10stake")
	//bech32AddressPrefix := types.AddrPrefixCfg{
	//	AccountAddr:   "cosmos",
	//	ValidatorAddr: "iva",
	//	ConsensusAddr: "ica",
	//	AccountPub:    "iap",
	//	ValidatorPub:  "ivp",
	//	ConsensusPub:  "icp",
	//}
	options := []types.Option{
		types.KeyDAOOption(store.NewMemory(nil)),
		types.TimeoutOption(30),
		types.KeyManagerOption(crypto.NewKeyManager()),
		types.BIP44PathOption(""),
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
