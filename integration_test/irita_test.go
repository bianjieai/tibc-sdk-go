package integration

import (
	"fmt"
	"testing"

	"github.com/irisnet/core-sdk-go/types"
)

const (
	nodeURIIritaA  = "tcp://localhost:26657"
	grpcAddrIritaA = "localhost:9090"
	chainIDIritaA  = "testA"
	keyNameIritaA  = "chainANode0"
	passwordIritaA = "12345678"
	keyStoreIritaA = `-----BEGIN TENDERMINT PRIVATE KEY-----
type: sm2
kdf: bcrypt
salt: DB87DBA0C7597ADAEB376421A8F6274E

kOMsXT79I5Jtja1F6pJd3QswWRk/AliSDBkwTCzmA3FVsbrNNnG6dwtP3u35YtGq
qSMVAD5ESuTSmSZPo4y3UqKwV8UbUnE5PxN9MMw=
=388z
-----END TENDERMINT PRIVATE KEY-----`
	chainLightClientNameA = "testCreateClientA"
	addressIritaA         = "iaa1rv9v0vrczgvgc43vygw88nxgwslqzlyf9ev55v"
)
const (
	nodeURIIritaC  = "tcp://localhost:36657"
	grpcAddrIritaC = "localhost:9190"
	chainIDIritaC  = "testC"
	keyNameIritaC  = "chainCNode0"
	passwordIritaC = "12345678"
	keyStoreIritaC = `-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: 085CC39C43FEF806B4C9615185C32EE3
type: sm2

50bD6y7XvPcSXJcuzJ/U8ZZPAZXrGCFBheCPfgJTCRaDhTalkToAO2Xd+uFnLrn7
g+SA5uqSDWildWyJMib55Q/p7pHyLRvFF+bDGcs=
=Qlej
-----END TENDERMINT PRIVATE KEY-----`
	chainLightClientNameC = "testCreateClientC"
	addressIritaC         = "iaa1xv92hhgng23547xrqy2d9e8pxqqurwd9cgdn05"
)

func Test_Irita(t *testing.T) {
	clientA, err := getIntegrationClient(nodeURIIritaA, grpcAddrIritaA, chainIDIritaA, keyNameIritaA, passwordIritaA, keyStoreIritaA, chainLightClientNameA)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error(), clientA)
		return
	}
	clientC, err := getIntegrationClient(nodeURIIritaC, grpcAddrIritaC, chainIDIritaC, keyNameIritaC, passwordIritaC, keyStoreIritaC, chainLightClientNameC)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error(), clientC)
		return
	}
	getTendermintjson(clientA.Tendermint, 100)
	getTendermintjson(clientC.Tendermint, 100)

	//sendTest(clientA, keyNameIritaA)
	//updateIritaClient(clientA, clientC)
}
func updateIritaClient(clientA, clientC Client) {
	updatetendetmintclientTest(clientA, clientC, chainLightClientNameC, keyNameIritaA)
	updatetendetmintclientTest(clientC, clientA, chainLightClientNameA, keyNameIritaC)
}

func sendTest(client Client, keyName string) {
	baseTx := types.BaseTx{
		From:               keyName,
		Gas:                300000,
		Memo:               "TEST",
		Mode:               types.Commit,
		Password:           "12345678",
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}

	dec := types.NewDec(100)

	decoin := types.DecCoin{
		Denom:  "upoint",
		Amount: dec,
	}
	//public, address, err222 := client.Find(keyName, "12345678")
	//fmt.Println(public, "\n", address, "\n", err222)
	amount := types.NewDecCoins(decoin)
	res, err := client.Bank.Send("iaa19g07ent7gq7fjm7xhm8eraxqx3gtrqy7w2gkuf", amount, baseTx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

}
