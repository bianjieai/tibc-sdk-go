package integration

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	tibc "github.com/bianjieai/tibc-sdk-go"
	tibcclient "github.com/bianjieai/tibc-sdk-go/client"
	"github.com/bianjieai/tibc-sdk-go/tendermint"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types"
	tenderminttypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

func getClientState(client tibc.Client, clientName string) {
	clientState1, err := client.GetClientState(clientName)
	if err != nil {
		panic(err)
	}
	fmt.Println(clientState1.String())
}
func getClientStates(client tibc.Client) {
	clientState1, err := client.GetClientStates()
	if err != nil {
		panic(err)
	}
	for _, value := range clientState1 {
		if value == nil {
			break
		}
		fmt.Println(value.String())
	}
}
func getheader(client Client, height int64, trustHeight tibcclient.Height, clientState tibctypes.ClientState) {
	res, err := client.QueryBlock(height)
	if err != nil {
		fmt.Println("QueryBlock fail:  ", err)
	}
	tmHeader := res.Block.Header

	rescommit, err := client.Commit(context.Background(), &res.BlockResult.Height)
	commit := rescommit.Commit
	signedHeader := &tenderminttypes.SignedHeader{
		Header: tmHeader.ToProto(),
		Commit: commit.ToProto(),
	}
	fmtheader := &tendermint.Header{
		SignedHeader:      signedHeader,
		ValidatorSet:      queryValidatorSet(height, client.Tendermint),
		TrustedHeight:     trustHeight,
		TrustedValidators: queryValidatorSet(int64(clientState.GetLatestHeight().GetRevisionHeight()), client.Tendermint),
	}
	b0, err := client.Tendermint.Codec.MarshalJSON(fmtheader)
	if err != nil {
		panic(err)
	}
	b0 = []byte(TenStaType + string(b0)[1:])
	clientStateName := tmHeader.ChainID + "_client_header.json"
	err = ioutil.WriteFile(clientStateName, b0, os.ModeAppend)
}

func getConsensusState(client tibc.Client, clientName string, height uint64) {
	consensusState1, err := client.GetConsensusState(clientName, height)
	if err != nil {
		panic(err)
	}
	fmt.Println(consensusState1.String())

}
func getConsensusStates(client tibc.Client) {
	consensusState1, err := client.GetConsensusStates("testCreateClient")
	if err != nil {
		panic(err)
	}
	fmt.Println("consensusState: ")
	for _, value := range consensusState1 {
		if value == nil {
			break
		}
		fmt.Println(value.String())
	}
}

func updateRinkebyEthClientTest(sourceClient Client, chainName, keyName, url string) {
	baseTx := types.BaseTx{
		From:               keyName,
		Gas:                0,
		Memo:               "TEST",
		Mode:               types.Commit,
		Password:           "12345678",
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}
	lightClientState, err := sourceClient.Tendermint.GetClientState(chainName)
	if err != nil {
		fmt.Println("GetClientState fail :", err, lightClientState)
		return
	}
	height := lightClientState.GetLatestHeight()
	ethHeader, err1 := GetRinkeyEthNodeHeader(url, height.GetRevisionHeight()+1)
	if err1 != nil {
		fmt.Println("GetEthNodeHeader fail :", err1, lightClientState)
		return
	}
	header := ethHeader.ToHeader()
	request := tibctypes.UpdateClientRequest{
		ChainName: chainName,
		Header:    &header,
	}
	fmt.Println("run : update client ", sourceClient.ChainName, ".", chainName, "start height : ", height, "want height:", height.GetRevisionHeight()+1)
	_, err = sourceClient.Tendermint.UpdateClient(request, baseTx)
	if err != nil {
		time.Sleep(time.Second * 5)
	}
	fmt.Println(" success : update client ", sourceClient.ChainName, ".", chainName, "end height : ", header.Height.RevisionHeight)
}

func updateEthClientTest(sourceClient Client, chainName, keyName, url string) {
	baseTx := types.BaseTx{
		From:               keyName,
		Gas:                0,
		Memo:               "TEST",
		Mode:               types.Commit,
		Password:           "12345678",
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}
	lightClientState, err := sourceClient.Tendermint.GetClientState(chainName)
	if err != nil {
		fmt.Println("GetClientState fail :", err, lightClientState)
		return
	}
	height := lightClientState.GetLatestHeight()
	ethHeader, err1 := GetEthNodeHeader(url, height.GetRevisionHeight()+1)
	if err1 != nil {
		fmt.Println("GetEthNodeHeader fail :", err1, lightClientState)
		return
	}
	header := ethHeader.ToHeader()
	request := tibctypes.UpdateClientRequest{
		ChainName: chainName,
		Header:    &header,
	}
	wantHeight := height.GetRevisionHeight() + 1
	fmt.Println("run : update client ", sourceClient.ChainName, ".", chainName, "start height : ", height, "want height:", wantHeight)
	_, err = sourceClient.Tendermint.UpdateClient(request, baseTx)
	if err != nil {
		time.Sleep(time.Second * 6)
		lightClientState, _ = sourceClient.Tendermint.GetClientState(chainName)
		height = lightClientState.GetLatestHeight()
		if height.GetRevisionHeight() != wantHeight {
			fmt.Println("update filed")
			return
		}
	}
	fmt.Println(" success : update client ", sourceClient.ChainName, ".", chainName, "end height : ", header.Height.RevisionHeight)
}

//destClient tibc.Client,
func updatetendetmintclientTest(sourceClient Client, destClient Client, chainName, keyName string) {
	baseTx := types.BaseTx{
		From:               keyName,
		Gas:                300000,
		Memo:               "TEST",
		Mode:               types.Commit,
		Password:           "12345678",
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}
	lightClientState, err := sourceClient.Tendermint.GetClientState(chainName)
	if err != nil {
		fmt.Println("UpdateClient fail :", err)
		return
	}
	height := int64(lightClientState.GetLatestHeight().GetRevisionHeight())
	stat, err1 := destClient.Status(context.Background())
	if err1 != nil {
		fmt.Println("get Status fail :", err1)
		return
	}
	height1 := stat.SyncInfo.LatestBlockHeight
	request := tibctypes.UpdateClientRequest{
		ChainName: chainName,
		Header: CreateTenderrmintHeader(
			destClient,
			height1,
			tibcclient.NewHeight(lightClientState.GetLatestHeight().GetRevisionNumber(), uint64(height)),
			lightClientState,
		),
	}
	fmt.Println("run : update client ", sourceClient.ChainName, ".", destClient.ChainName)
	_, err = sourceClient.Tendermint.UpdateClient(request, baseTx)
	if err != nil {
		fmt.Println("UpdateClient fail :", err)
		return
	}
	fmt.Println(" success : update client ", sourceClient.ChainName, ".", destClient.ChainName)
}

func CreateTenderrmintHeader(client Client, height int64, trustHeight tibcclient.Height, clientState tibctypes.ClientState) *tendermint.Header {
	res, err := client.QueryBlock(height)
	if err != nil {
		fmt.Println("QueryBlock fail:  ", err)
	}
	tmHeader := res.Block.Header

	rescommit, err := client.Commit(context.Background(), &res.BlockResult.Height)
	commit := rescommit.Commit
	signedHeader := &tenderminttypes.SignedHeader{
		Header: tmHeader.ToProto(),
		Commit: commit.ToProto(),
	}

	return &tendermint.Header{
		SignedHeader:      signedHeader,
		ValidatorSet:      queryValidatorSet(height, client.Tendermint),
		TrustedHeight:     trustHeight,
		TrustedValidators: queryValidatorSet(int64(clientState.GetLatestHeight().GetRevisionHeight()), client.Tendermint),
	}

}

func queryValidatorSet(height int64, client tibc.Client) *tenderminttypes.ValidatorSet {
	page := 1
	prepage := 200
	validators, err := client.Validators(context.Background(), &height, &page, &prepage)
	if err != nil {
		fmt.Println("queryValidatorSet fail :", err)
	}
	validatorSet, err := tmtypes.NewValidatorSet(validators.Validators[:]).ToProto()
	if err != nil {
		fmt.Println("queryValidatorSet fail :", err)
	}
	return validatorSet
}

func updatebscclientTest(sourceClient Client, destchainUrl, chainName, keyname string) {
	baseTx := types.BaseTx{
		From:               keyname,
		Gas:                0,
		Memo:               "TEST",
		Mode:               types.Commit,
		Password:           "12345678",
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}
	lightClientState, err := sourceClient.Tendermint.GetClientState(chainName)
	if err != nil {
		fmt.Println("GetClientState fail :", err, lightClientState)
		return
	}
	height := lightClientState.GetLatestHeight()
	bscHeader, err1 := GetNodeHeader(destchainUrl, height.GetRevisionHeight()+1)
	if err1 != nil {
		fmt.Println("GetClientState fail :", err1, lightClientState)
		return
	}
	header := bscHeader.ToHeader()
	request := tibctypes.UpdateClientRequest{
		ChainName: chainName,
		Header:    &header,
	}
	fmt.Println("run : update client ", sourceClient.ChainName, ".", chainName, "start height : ", header.Height.RevisionHeight)
	_, err = sourceClient.Tendermint.UpdateClient(request, baseTx)
	if err != nil {
		fmt.Println("UpdateClient fail :", err)
		return
	}
	fmt.Println(" success : update client ", sourceClient.ChainName, ".", chainName, "end height : ", header.Height.RevisionHeight+1)

}
