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
salt: 3823352403F3B83C71DC32A04BC72B54
type: secp256k1

d9EsZANuTcqPETSFbK1A8leVZopyP4hJ58iojC4Ob8tkp07EgqJuDps52zIEaJpm
iBfDOTLg1REK42bKoRs1+uJ2Bl3Z8/HkkwWdmSo=
=VIyW
-----END TENDERMINT PRIVATE KEY-----`
	chainALightClientName = "testCreateClientA"
	addressA              = "cosmos1h07pjtuszwz5jt0g62lmlx86pk80wvw8heqhhm"
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
