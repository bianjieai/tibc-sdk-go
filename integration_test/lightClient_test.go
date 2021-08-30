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
salt: D145AFB29D332974833DB1AD5912D3D1
type: secp256k1

KCtCKbIq/EVtNGdZSoa4cwkLfMk9Z7tMQ6P1DrSe0KqkmMutzA2hh3WbtrSdJsgD
KZM9JwHI4ROYyXGsTD5s8oL+kjErG7zcpjhbX7g=
=K8L7
-----END TENDERMINT PRIVATE KEY-----`
	chainALightClientName = "testCreateClientA"
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
	chainBLightClientName = "testCreateClientB"
)
const (
	nodeURI2  = "tcp://192.168.232.140:26657"
	grpcAddr2 = "192.168.232.140:9090"
	chainID2  = "testC"
	keyName2  = "chainCNode0"
	password2 = "12345678"
	keyStore2 = `-----BEGIN TENDERMINT PRIVATE KEY-----
salt: CA512713FF0FA3BCDFA68F2CEE202789
type: secp256k1
kdf: bcrypt

mROEY8pNQewR8cAv1MrypPmM5V/iyAMWZUWL8h7Zvs0bdtbZSdNLqGm0Mtw2x3b8
8MfdI5iQkuSAkwWs33lwNs69l72mtIH76dV2KVU=
=PrCi
-----END TENDERMINT PRIVATE KEY-----`
	chainCLightClientName = "testCreateClientC"
)

func Test_integrationClient(t *testing.T) {
	clientA, err := getIntegrationClient(nodeURI0, grpcAddr0, chainID0, keyName0, password0, keyStore0, chainALightClientName)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error())
		return
	}
	clientB, err := getIntegrationClient(nodeURI1, grpcAddr1, chainID1, keyName1, password1, keyStore1, chainBLightClientName)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error())
		return
	}
	clientC, err := getIntegrationClient(nodeURI2, grpcAddr2, chainID2, keyName2, password2, keyStore2, chainCLightClientName)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error())
		return
	}
	updateAllCient(clientA, clientB, clientC)
	//cleanPacket(clientA,clientB,1,keyName0)
	//recvCleanPacket(clientA,clientB,keyName1,"1ECE3853D71E786198CD7241BF774E281BFB5DD1CDF3704FE8C4ADCB0E400DC6")

	//single jump A to B then return
	nftAtoB(clientA, clientB, clientC)
	nftBReturntoA(clientA, clientB, clientC)

	//double jump B to C  (relayer A) then return
	nftBtoC(clientA, clientB, clientC)
	nftCReturntoB(clientA, clientB, clientC)
}

func nftAtoB(clientA, clientB, clientC Client) {
	fmt.Println("testnftTransfer: (A to B)")
	txhash, err := nftTransfer(clientA, keyName0, "atobtestclass", "atobtestid", "iaa1swva40c0js3xfc6dcqgeml34kyzqcmpdnhqks6", chainBLightClientName, "")
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
	txhash, err := nftTransfer(clientB, keyName1, "tibc/nft/testCreateClientA/atobtestclass", "atobtestid", "iaa19gqlam0dq3hda9k2guc64a2r2s9le0eysugsj7", chainALightClientName, "")
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
	updateclientTest(clientA, clientB, chainBLightClientName, keyName0)
	updateclientTest(clientA, clientC, chainCLightClientName, keyName0)
	updateclientTest(clientB, clientA, chainALightClientName, keyName1)
	updateclientTest(clientC, clientA, chainALightClientName, keyName2)
}
func getIntegrationClient(nodeURI, grpcAddr, chainID, keyName, password, keyStore, chainName string) (Client, tibctypes.IError) {
	feeCoin, err := types.ParseDecCoins("10stake")
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
	txhash, err := nftTransfer(clientB, keyName1, "btoctestclass", "btoctestid", "iaa1x0c9ycm0hywjzqvwy4xlepzrp3sph4z88mdzsw", chainCLightClientName, chainALightClientName)
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
	txhash, err := nftTransfer(clientC, keyName2, "tibc/nft/testCreateClientB/btoctestclass", "btoctestid", "iaa1swva40c0js3xfc6dcqgeml34kyzqcmpdnhqks6", chainBLightClientName, chainALightClientName)
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
