package integration

import (
	"fmt"
	"testing"
	"time"

	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	"github.com/irisnet/core-sdk-go/common/crypto"
	"github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/store"
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
	nodeURI2  = "tcp://192.168.232.140:26657"
	grpcAddr2 = "192.168.232.140:9090"
	chainID2  = "testC"
	keyName2  = "chainCNode0"
	password2 = "12345678"
	keyStore2 = `-----BEGIN TENDERMINT PRIVATE KEY-----
type: secp256k1
kdf: bcrypt
salt: 202B62B4E71F2EEC9B8386CAEC2B5BD2

Reo/ka7X9Va2hySLE4yqoOx0jjZn9LILs1mXCuqpqM/tTJm02oDxiQ7B+xbEOSbr
0nAstU3hEsYEreMg7ihO+7dK0ufXEx8JEbnTIHU=
=wiwf
-----END TENDERMINT PRIVATE KEY-----`
	chainCLightClientName = "testCreateClientC"
	addressC              = "iaa10nfdefym9vg7c288fm4790833ee5f4p0g8w3ej"
)

func Test_integrationClientTen(t *testing.T) {
	//clientA, err := getIntegrationClient(nodeURI0, grpcAddr0, chainID0, keyName0, password0, keyStore0, chainALightClientName)
	//if err != nil {
	//	fmt.Println(err.Codespace(), err.Code(), err.Error())
	//	return
	//}
	//clientB, err := getIntegrationClient(nodeURI1, grpcAddr1, chainID1, keyName1, password1, keyStore1, chainBLightClientName)
	//if err != nil {
	//	fmt.Println(err.Codespace(), err.Code(), err.Error())
	//	return
	//}
	clientC, err := getIntegrationClient(nodeURI2, grpcAddr2, chainID2, keyName2, password2, keyStore2, chainCLightClientName)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error())
		return
	}
	for i := 1; i <= 300; i++ {
		updateEthClientTest(clientC, "ethclient", keyName2)
		fmt.Println("    seq : ", i)
	}

	//updateAllCient(clientA, clientB, clientC)
	//getETHjson(clientA.Tendermint)
	//cleanPacket(clientA,clientB,1,keyName0)
	//recvCleanPacket(clientA,clientB,keyName1,"1ECE3853D71E786198CD7241BF774E281BFB5DD1CDF3704FE8C4ADCB0E400DC6")

	//single jump A to B then return
	//nftAtoB(clientA, clientB, clientC)
	//nftBReturntoA(clientA, clientB, clientC)

	//double jump B to C  (relayer A) then return
	//nftBtoC(clientA, clientB, clientC)
	//nftCReturntoB(clientA, clientB, clientC)

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

func nftAtoB(clientA, clientB, clientC Client) {
	fmt.Println("testnftTransfer: (A to B)")
	txhash, err := nftTransfer(clientA, keyName0, "atobtestclass", "atobtestid", addressB, chainBLightClientName, "")
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error())
		return
	}
	fmt.Println("txhash : ", txhash)
	time.Sleep(time.Second * 5)
	updateAllCient(clientA, clientB, clientC)

	fmt.Println("recv packet")
	txhash, err = packetRecive(clientA, clientB, keyName1, txhash)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error())
		return
	}
	fmt.Println("txhash : ", txhash)
	time.Sleep(time.Second * 5)
	updateAllCient(clientA, clientB, clientC)

	fmt.Println("send ack : ")
	txhash, err = sendAck(clientA, clientB, keyName0, txhash)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("txhash : ", txhash)
}
func nftBReturntoA(clientA, clientB, clientC Client) {
	fmt.Println("testnftTransfer: (nft B Return to A)")
	txhash, err := nftTransfer(clientB, keyName1, "tibc/nft/testCreateClientA/atobtestclass", "atobtestid", addressA, chainALightClientName, "")
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error())
		return
	}
	fmt.Println("txhash : ", txhash)
	time.Sleep(time.Second * 5)
	updateAllCient(clientA, clientB, clientC)

	fmt.Println("recv packet")
	txhash, err = packetRecive(clientB, clientA, keyName0, txhash)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error())
		return
	}
	fmt.Println("txhash : ", txhash)
	time.Sleep(time.Second * 5)
	updateAllCient(clientA, clientB, clientC)

	fmt.Println("send ack : ")
	txhash, err = sendAck(clientB, clientA, keyName1, txhash)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("txhash : ", txhash)
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
		types.TimeoutOption(10),
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

func nftBtoC(clientA, clientB, clientC Client) {
	fmt.Println("testnftTransfer: (B to C)")
	txhash, err := nftTransfer(clientB, keyName1, "btoctestclass", "btoctestid", addressC, chainCLightClientName, chainALightClientName)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error())
		return
	}
	fmt.Println("txhash : ", txhash)
	time.Sleep(time.Second * 5)
	updateAllCient(clientA, clientB, clientC)

	fmt.Println("recv packet1")
	txhash, err = packetRecive(clientB, clientA, keyName0, txhash)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error())
		return
	}
	fmt.Println("txhash : ", txhash)
	time.Sleep(time.Second * 5)
	updateAllCient(clientA, clientB, clientC)

	fmt.Println("recv packet2")
	txhash, err = packetRecive(clientA, clientC, keyName2, txhash)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error())
		return
	}
	fmt.Println("txhash : ", txhash)
	time.Sleep(time.Second * 5)
	updateAllCient(clientA, clientB, clientC)

	fmt.Println("send ack1 : ")
	txhash, err = sendAck(clientA, clientC, keyName0, txhash)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("txhash : ", txhash)
	time.Sleep(time.Second * 5)
	updateAllCient(clientA, clientB, clientC)

	fmt.Println("send ack2 : ")
	txhash, err = sendAck(clientB, clientA, keyName1, txhash)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("txhash : ", txhash)
}

func nftCReturntoB(clientA, clientB, clientC Client) {

	fmt.Println("testnftTransfer: (C to B)")
	txhash, err := nftTransfer(clientC, keyName2, "tibc/nft/testCreateClientB/btoctestclass", "btoctestid", addressB, chainBLightClientName, chainALightClientName)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error())
		return
	}
	fmt.Println("txhash : ", txhash)
	time.Sleep(time.Second * 5)
	updateAllCient(clientA, clientB, clientC)

	fmt.Println("recv packet1")
	txhash, err = packetRecive(clientC, clientA, keyName0, txhash)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error())
		return
	}
	fmt.Println("txhash : ", txhash)
	time.Sleep(time.Second * 5)
	updateAllCient(clientA, clientB, clientC)

	fmt.Println("recv packet2")
	txhash, err = packetRecive(clientA, clientB, keyName1, txhash)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error())
		return
	}
	fmt.Println("txhash : ", txhash)
	time.Sleep(time.Second * 5)
	updateAllCient(clientA, clientB, clientC)

	fmt.Println("send ack1 : ")
	txhash, err = sendAck(clientA, clientB, keyName0, txhash)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("txhash : ", txhash)
	time.Sleep(time.Second * 5)
	updateAllCient(clientA, clientB, clientC)

	fmt.Println("send ack2 : ")
	txhash, err = sendAck(clientC, clientA, keyName2, txhash)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("txhash : ", txhash)

}
