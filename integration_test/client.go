package integration

import (
	"github.com/bianjieai/tibc-sdk-go"
	tibcclient "github.com/bianjieai/tibc-sdk-go/client"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	sdk "github.com/irisnet/core-sdk-go"
	"github.com/irisnet/core-sdk-go/bank"
	"github.com/irisnet/core-sdk-go/client"
	keys "github.com/irisnet/core-sdk-go/client"
	commoncodec "github.com/irisnet/core-sdk-go/common/codec"
	cryptotypes "github.com/irisnet/core-sdk-go/common/codec/types"
	"github.com/irisnet/core-sdk-go/types"
	txtypes "github.com/irisnet/core-sdk-go/types/tx"
	"github.com/tendermint/tendermint/libs/log"
)

type clientforlightclient struct {
	logger         log.Logger
	moduleManager  map[string]types.Module
	encodingConfig types.EncodingConfig

	types.BaseClient
	Bank             bank.Client
	Key              keys.Client
	TendermintClient tibcclient.ChainClient
}

func newClient(cfg types.ClientConfig) clientforlightclient {
	encodingConfig := makeEncodingConfig()

	// create a instance of baseClient
	baseClient := client.NewBaseClient(cfg, encodingConfig, nil)
	bankClient := bank.NewClient(baseClient, encodingConfig.Marshaler)
	tendermintClient := tibc_sdk_go.NewClient(baseClient, encodingConfig.Marshaler)
	keysClient := keys.NewKeysClient(cfg, baseClient)

	client := &clientforlightclient{
		logger:           baseClient.Logger(),
		BaseClient:       baseClient,
		moduleManager:    make(map[string]types.Module),
		encodingConfig:   encodingConfig,
		Bank:             bankClient,
		TendermintClient: tendermintClient,
		Key:              keysClient,
	}

	client.registerModule(
		bankClient,
		tendermintClient,
	)

	return *client
}
func (client clientforlightclient) registerModule(ms ...types.Module) {
	for _, m := range ms {
		m.RegisterInterfaceTypes(client.encodingConfig.InterfaceRegistry)
	}
}

func makeEncodingConfig() types.EncodingConfig {
	amino := commoncodec.NewLegacyAmino()
	interfaceRegistry := cryptotypes.NewInterfaceRegistry()
	marshaler := commoncodec.NewProtoCodec(interfaceRegistry)
	txCfg := txtypes.NewTxConfig(marshaler, txtypes.DefaultSignModes)

	encodingConfig := types.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}
	sdk.RegisterLegacyAminoCodec(encodingConfig.Amino)
	sdk.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	tibctypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
