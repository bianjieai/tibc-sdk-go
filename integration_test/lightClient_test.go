package integration

import (
	"fmt"
	"testing"

	"github.com/irisnet/core-sdk-go/common/crypto"
	"github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/store"
)

const (
	nodeURI  = "tcp://192.168.232.140:26657"
	grpcAddr = "192.168.232.140:9090"
	chainID  = "testnet"
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

	client := newClient(cfg)

	if err != nil {
		panic(err)
	}
	_, err = client.Key.Import(keyName, password, keyStore)
	if err != nil {
		panic(err)
	}
	getClientState(client)
	getconesState(client)
	getConsensusState(client)
	getConsensusState(client)
}

func getClientState(client clientforlightclient) {
	clientState1, err := client.TendermintClient.GetClientState("testCreateClient")
	if err != nil {
		panic(err)
	}
	fmt.Println(clientState1.String())
}
func getconesState(client clientforlightclient) {
	clientState1, err := client.TendermintClient.GetClientStates()
	if err != nil {
		panic(err)
	}
	for _,value := range clientState1 {
		if value == nil {
			break
		}
		fmt.Println(value.String())
	}
}

func getConsensusState(client clientforlightclient) {
	consensusState1, err := client.TendermintClient.GetConsensusState("testCreateClient",8)
	if err != nil {
		panic(err)
	}

	fmt.Println(consensusState1.String())

}
func getConsensusStates(client clientforlightclient) {
	consensusState1, err := client.TendermintClient.GetConsensusStates("testCreateClient")
	if err != nil {
		panic(err)
	}
	for _,value := range consensusState1 {
		if value == nil {
			break
		}
		fmt.Println(value.String())
	}

}
