package client

import (
	"fmt"
	sdk "github.com/irisnet/core-sdk-go"
	"github.com/irisnet/core-sdk-go/common/crypto"
	"github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/store"
	"testing"
)

const (
	nodeURI  = "tcp://localhost:26657"
	grpcAddr = "localhost:9090"
	chainID  = "irita-test"
	keyName  = "node0"
	password = "12345678"
	keyStore = `-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: 618CFFE14C9AABCD4C6E1F4504FDFBDE
type: secp256k1

pL3QGEQaLve37wBPAPEz9hSP3ZZeZeX/j46hFOuOsrNtOYUJsTIah2UODmIyF5NY
mGlFzRtegvA4+A8n5sCIjwqnPMz8hX0IF9ltLjQ=
=rDPe
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
	bech32AddressPrefix := types.AddrPrefixCfg{
		AccountAddr:   "iaa",
		ValidatorAddr: "iva",
		ConsensusAddr: "ica",
		AccountPub:    "iap",
		ValidatorPub:  "ivp",
		ConsensusPub:  "icp",
	}

	feeCoin, err := types.ParseDecCoins("10stake")
	options := []types.Option{
		types.KeyDAOOption(store.NewMemory(nil)),
		types.TimeoutOption(10),
		types.TokenManagerOption(TokenManager{}),
		types.KeyManagerOption(crypto.NewKeyManager()),
		types.Bech32AddressPrefixOption(bech32AddressPrefix),
		types.BIP44PathOption(""),
		types.FeeOption(feeCoin),
	}

	cfg, err := types.NewClientConfig(nodeURI, grpcAddr, chainID, options...)
	if err != nil {
		panic(err)
	}

	client := NewClient(cfg)

	if err != nil {
		panic(err)
	}
	//fmt.Println(client)
	//addr , err := client.Key.Recover("test", "12345678", "strategy project text close add advance hint gaze rent future shoe winner dust reform scrub diagram trash ring critic vault edge potato pyramid fee")
	_, err = client.Key.Import(keyName, password, keyStore)
	//a,err:=client.BaseClient.QueryAccount("iaa1vmhh0lekt7nha9gkngqdjwleqvn9m24aemp7g2")
	if err != nil {
		panic(err)
	}
	//fmt.Println(client.sdkClient.Bank.QueryAccount(addrse))
	//bankSend(addr)
	//fmt.Println(client.sdkClient.GenConn())
	//fmt.Println(client.tendermintClient.Getconnn())
	getClientState(client)
	//fmt.Println(a)
}

func getClientState(client clientforlightclient) {
	clientState1, err := client.TendermintClient.GetClientState("testCreateClient")
	if err != nil {
		panic(err)
	}
	fmt.Println(clientState1.String())

}

func bankSend(client sdk.Client) {
	addr, err := client.Key.Recover("test", "12345678", "strategy project text close add advance hint gaze rent future shoe winner dust reform scrub diagram trash ring critic vault edge potato pyramid fee")
	//a, err := client.Key.Import("v1", "12345678", "-----BEGIN TENDERMINT PRIVATE KEY-----\nkdf: bcrypt\nsalt: E4EEAD39E485366E68E70812E16081D1\ntype: sm2\n\nXieOvCObvTx0I+7WIYlUeYIgrZvIf4aZFUaDCrWSPCIPmPHbo5fLSdax2vuCQfXl\nMEA7MIdFgxzKss4M/cmnwZOPoOjz/jZKPaE3q8g=\n=H9E3\n-----END TENDERMINT PRIVATE KEY-----")
	//a,err:=client.BaseClient.QueryAccount("iaa1vmhh0lekt7nha9gkngqdjwleqvn9m24aemp7g2")
	if err != nil {
		panic(err)
	}
	acc, err := client.Bank.QueryAccount(addr)
	fmt.Println(acc.Coins.String())
	coins, _ := types.ParseDecCoins("100upoint")
	to := "iaa18xqa8yd57jfvh0cdkex7hqn6t7zspzwwl6p7e3"
	baseTx := types.BaseTx{
		From:               "test",
		Gas:                400000,
		Memo:               "TEST",
		Mode:               types.Commit,
		Password:           "12345678",
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}

	res, err := client.Bank.Send(to, coins, baseTx)
	//res, err :=client.QueryToken("upoint")

	if err != nil {
		panic(err)
	}
	fmt.Println(res)

}
