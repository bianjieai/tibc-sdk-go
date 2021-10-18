package integration

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/irisnet/core-sdk-go/types"
)

const (
	nodeURIIritaA  = "tcp://localhost:26657"
	grpcAddrIritaA = "localhost:9090"
	chainIDIritaA  = "testA"
	keyNameIritaA  = "chainANode0"
	passwordIritaA = "12345678"
	keyStoreIritaA = `-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: E8B7448668D583E92FB5BA4F80A3F86C
type: sm2

GMOWF0k8nOuSEMNrAkEXIlWk9vc2RLFzr3pMjWJSg5hlpuzUEPePVIm2m4BCNUD6
vKXJvRcAtRCpFWLlTcl1UdOw4OXWejmP/Jt4QEw=
=vvsl
-----END TENDERMINT PRIVATE KEY-----`
	chainLightClientNameA = "testCreateClientA"
	addressIritaA         = "iaa1nu7awd4evcn4ts4h0su735nhv9vfejkkv3unpc"
)
const (
	nodeURIIritaC  = "tcp://localhost:36657"
	grpcAddrIritaC = "localhost:9190"
	chainIDIritaC  = "testC"
	keyNameIritaC  = "chainCNode0"
	passwordIritaC = "12345678"
	keyStoreIritaC = `-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: EFD047FF1911092D736FBA56C6A6FFF9
type: sm2

RzaJ5LV3qZ6bO/xf6AiIPAmpop4t0roG1y+3GWBWPw4PAsSeN5JajSgyadlad1C7
aGm0ZlxzZL2U7DxTk5neFiNePMz7H4U25JDHpf4=
=DcPH
-----END TENDERMINT PRIVATE KEY-----`
	chainLightClientNameC = "testCreateClientC"
	addressIritaC         = "iaa1v8a9a3pql7fznez0yg8tlg03huh34a4u9u04h7"
)

func Test_Irita(t *testing.T) {
	clientA, err := getIntegrationClient(nodeURIIritaA, grpcAddrIritaA, chainIDIritaA, keyNameIritaA, passwordIritaA, keyStoreIritaA, chainLightClientNameA)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error(), clientA)
		return
	}
	address, err2 := clientA.QueryAddress(keyNameIritaA, passwordIritaA)
	if err2 != nil {
		return
	}
	require.Equal(t, address.String(), addressIritaA)

	clientC, err := getIntegrationClient(nodeURIIritaC, grpcAddrIritaC, chainIDIritaC, keyNameIritaC, passwordIritaC, keyStoreIritaC, chainLightClientNameC)
	if err != nil {
		fmt.Println(err.Codespace(), err.Code(), err.Error(), clientC)
		return
	}
	address, err2 = clientC.QueryAddress(keyNameIritaC, passwordIritaC)
	if err2 != nil {
		return
	}
	require.Equal(t, address.String(), addressIritaC)

	getTendermintjson(clientA.Tendermint, 10)
	getTendermintjson(clientC.Tendermint, 10)
	sendTest(clientA, keyNameIritaA)
	updateIritaClient(clientA, clientC)
	res, err := packetRecive(clientC, clientA, keyNameIritaA, "8200D3BC9CE7244BF33088C524ABFEBF96759E7AB30DE3013AAE92BC0579A72F")
	require.Nil(t, err)
	fmt.Println(res)
	updateIritaClient(clientA, clientC)
	updateIritaClient(clientA, clientC)
	res, err = sendAck(clientC, clientA, keyNameIritaC, "614FD3941370E7405CD89BABE9312D093743FB3F00CCFF304A42771194A73B6A")
	require.Nil(t, err)
	fmt.Println(res)
	updateIritaClient(clientA, clientC)

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
