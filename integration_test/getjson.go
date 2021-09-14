package integration

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	tibceth "github.com/bianjieai/tibc-sdk-go/eth"

	tibc "github.com/bianjieai/tibc-sdk-go"
	tibcbsc "github.com/bianjieai/tibc-sdk-go/bsc"
	tibcclient "github.com/bianjieai/tibc-sdk-go/client"
	"github.com/bianjieai/tibc-sdk-go/commitment"
	"github.com/bianjieai/tibc-sdk-go/tendermint"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

const (
	TenConType = "{\"@type\":\"/tibc.lightclients.tendermint.v1.ConsensusState\","
	TenStaType = "{\"@type\":\"/tibc.lightclients.tendermint.v1.ClientState\","
	BscConType = "{\"@type\":\"/tibc.lightclients.bsc.v1.ConsensusState\","
	BscStaType = "{\"@type\":\"/tibc.lightclients.bsc.v1.ClientState\","
	EthConType = "{\"@type\":\"/tibc.lightclients.eth.v1.ConsensusState\","
	EthStaType = "{\"@type\":\"/tibc.lightclients.eth.v1.ClientState\","
)

const (
	ethurl         = "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"
	mainurl        = "https://bsc-dataseed1.binance.org"
	testneturl     = "https://data-seed-prebsc-1-s1.binance.org:8545/"
	epoch          = uint64(200)
	TestnetChainId = 97
)

//Generate a JSON file needed to create the light client
//Add the following string to the header after the file is generated
//"@type":"/tibc.lightclients.tendermint.v1.ClientState",
//"@type":"/tibc.lightclients.tendermint.v1.ConsensusState",
func getTendermintjson(client tibc.Client, height int64) {

	//ClientState
	var fra = tendermint.Fraction{
		Numerator:   1,
		Denominator: 3,
	}
	res, err := client.QueryBlock(height)
	if err != nil {
		fmt.Println("QueryBlock fail:  ", err)
	}
	tmHeader := res.Block.Header
	lastHeight := tibcclient.NewHeight(0, 4)
	var clientstate = &tendermint.ClientState{
		ChainId:         tmHeader.ChainID,
		TrustLevel:      fra,
		TrustingPeriod:  time.Hour * 24 * 7 * 2,
		UnbondingPeriod: time.Hour * 24 * 7 * 3,
		MaxClockDrift:   time.Second * 10,
		LatestHeight:    lastHeight,
		ProofSpecs:      commitment.GetSDKSpecs(),
		MerklePrefix:    commitment.MerklePrefix{KeyPrefix: []byte("tibc")},
		TimeDelay:       0,
	}
	//ConsensusState
	var consensusState = &tendermint.ConsensusState{
		Timestamp:          tmHeader.Time,
		Root:               commitment.NewMerkleRoot([]byte("app_hash")),
		NextValidatorsHash: queryValidatorSet1(res.Block.Height, client).Hash(),
	}

	b0, err := client.Marshaler.MarshalJSON(clientstate)
	if err != nil {
		panic(err)
	}
	b0 = []byte(TenStaType + string(b0)[1:])
	clientStateName := tmHeader.ChainID + "_client_state.json"
	err = ioutil.WriteFile(clientStateName, b0, os.ModeAppend)
	if err != nil {
		return
	}
	b1, err := client.Marshaler.MarshalJSON(consensusState)
	if err != nil {
		panic(err)
	}
	b1 = []byte(TenConType + string(b1)[1:])
	clientConsensusStateName := tmHeader.ChainID + "_consensus_state.json"
	err = ioutil.WriteFile(clientConsensusStateName, b1, os.ModeAppend)
	if err != nil {
		return
	}
}

func queryValidatorSet1(height int64, client tibc.Client) *tmtypes.ValidatorSet {
	validators, err := client.Validators(context.Background(), &height, nil, nil)
	if err != nil {
		fmt.Println("queryValidatorSet1 fail :", err)
	}
	validatorSet := tmtypes.NewValidatorSet(validators.Validators)
	if err != nil {
		fmt.Println("queryValidatorSet fail :", err)
	}
	return validatorSet
}

type RestClient struct {
	Addr       string
	restClient *http.Client
}

func NewRestClient() *RestClient {
	return &RestClient{
		restClient: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   5,
				DisableKeepAlives:     false,
				IdleConnTimeout:       time.Second * 300,
				ResponseHeaderTimeout: time.Second * 300,
				TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
			},
			Timeout: time.Second * 300,
		},
	}
}
func getBSCjson(client tibc.Client) {
	rc := NewRestClient()
	height, err := GetBlockHeight(rc, testneturl)
	if err != nil {
		fmt.Println(err)
		return
	}
	genesisHeight := height - height%epoch - 2*epoch
	header, err := GetNodeHeader(rc, testneturl, genesisHeight)
	//validators, err := tibcbsc.ParseValidators(header.Extra)
	genesisValidatorHeader, err := GetNodeHeader(rc, testneturl, genesisHeight-epoch)
	genesisValidators, err := tibcbsc.ParseValidators(genesisValidatorHeader.Extra)
	number := tibcclient.NewHeight(0, header.Number.Uint64())
	clientState := tibcbsc.ClientState{
		Header:          header.ToHeader(),
		ChainId:         TestnetChainId,
		Epoch:           epoch,
		BlockInteval:    3,
		Validators:      genesisValidators,
		ContractAddress: []byte("0x00"),
		TrustingPeriod:  200,
	}

	consensusState := tibcbsc.ConsensusState{
		Timestamp: header.Time,
		Number:    number,
		Root:      header.Root[:],
	}
	b0, err := client.Marshaler.MarshalJSON(&clientState)
	if err != nil {
		panic(err)
	}
	b0 = []byte(BscStaType + string(b0)[1:])
	clientStateName := "bsc_client_state.json"
	err = ioutil.WriteFile(clientStateName, b0, os.ModeAppend)
	if err != nil {
		return
	}
	b1, err := client.Marshaler.MarshalJSON(&consensusState)
	if err != nil {
		panic(err)
	}
	b1 = []byte(BscConType + string(b1)[1:])
	clientConsensusStateName := "bsc_consensus_state.json"
	err = ioutil.WriteFile(clientConsensusStateName, b1, os.ModeAppend)
	if err != nil {
		return
	}
}

func getETHjson(client tibc.Client) {
	rc := NewRestClient()
	height, err := GetBlockHeight(rc, ethurl)
	if err != nil {
		fmt.Println(err)
		return
	}
	genesisHeight := height - height%epoch - 2*epoch
	header, err := GetEthNodeHeader(rc, ethurl, genesisHeight)
	//validators, err := tibcbsc.ParseValidators(header.Extra)
	number := tibcclient.NewHeight(0, header.Number.Uint64())
	clientState := tibceth.ClientState{
		Header:          header.ToHeader(),
		ChainId:         1,
		ContractAddress: []byte("0x00"),
		TrustingPeriod:  200,
		TimeDelay:       0,
		BlockDelay:      7,
	}

	consensusState := tibceth.ConsensusState{
		Timestamp: header.Time,
		Number:    number,
		Root:      header.Root[:],
		Header:    header.ToHeader(),
	}
	b0, err := client.Marshaler.MarshalJSON(&clientState)
	if err != nil {
		panic(err)
	}
	b0 = []byte(EthStaType + string(b0)[1:])
	clientStateName := "eth_client_state.json"
	err = ioutil.WriteFile(clientStateName, b0, os.ModeAppend)
	if err != nil {
		return
	}
	b1, err := client.Marshaler.MarshalJSON(&consensusState)
	if err != nil {
		panic(err)
	}
	b1 = []byte(EthConType + string(b1)[1:])
	clientConsensusStateName := "eth_consensus_state.json"
	err = ioutil.WriteFile(clientConsensusStateName, b1, os.ModeAppend)
	if err != nil {
		return
	}
}

type heightReq struct {
	JsonRpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	Id      uint     `json:"id"`
}

type heightRsp struct {
	JsonRpc string     `json:"jsonrpc"`
	Result  string     `json:"result,omitempty"`
	Error   *jsonError `json:"error,omitempty"`
	Id      uint       `json:"id"`
}

type jsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type blockReq struct {
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      uint          `json:"id"`
}
type blockRsp struct {
	JsonRPC string           `json:"jsonrpc"`
	Result  *ethtypes.Header `json:"result,omitempty"`
	Error   *jsonError       `json:"error,omitempty"`
	Id      uint             `json:"id"`
}

func GetBlockHeight(rc *RestClient, url string) (height uint64, err error) {

	req := &heightReq{
		JsonRpc: "2.0",
		Method:  "eth_blockNumber",
		Params:  make([]string, 0),
		Id:      1,
	}
	reqData, err := json.Marshal(req)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight: marshal req err: %s", err)
	}
	rspData, err := rc.SendRestRequest(url, reqData)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight err: %s", err)
	}

	rsp := &heightRsp{}
	err = json.Unmarshal(rspData, rsp)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight, unmarshal resp err: %s", err)
	}
	if rsp.Error != nil {
		return 0, fmt.Errorf("GetNodeHeight, unmarshal resp err: %s", rsp.Error.Message)
	}
	height, err = strconv.ParseUint(rsp.Result, 0, 64)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight, parse resp height %s failed", rsp.Result)
	} else {
		return height, nil
	}
}
func (self *RestClient) SendRestRequest(addr string, data []byte) ([]byte, error) {
	resp, err := self.restClient.Post(addr, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, fmt.Errorf("http post request:%s error:%s", data, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read rest response body error:%s", err)
	}
	return body, nil
}

func GetNodeHeader(restClient *RestClient, url string, height uint64) (*tibcbsc.BscHeader, error) {
	params := []interface{}{fmt.Sprintf("0x%x", height), true}
	req := &blockReq{
		JsonRpc: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  params,
		Id:      1,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetNodeHeight: marshal req err: %s", err)
	}
	rspdata, err := restClient.SendRestRequest(url, reqdata)
	if err != nil {
		return nil, fmt.Errorf("GetNodeHeight err: %s", err)
	}
	rsp := &blockRsp{}
	err = json.Unmarshal(rspdata, rsp)
	if err != nil {
		return nil, fmt.Errorf("GetNodeHeight, unmarshal resp err: %s", err)
	}
	if rsp.Error != nil {
		return nil, fmt.Errorf("GetNodeHeight, return error: %s", rsp.Error.Message)
	}

	if rsp.Result == nil {
		return nil, errors.New("GetNodeHeight, no result")
	}

	header := rsp.Result
	return &tibcbsc.BscHeader{
		ParentHash:  header.ParentHash,
		UncleHash:   header.UncleHash,
		Coinbase:    header.Coinbase,
		Root:        header.Root,
		TxHash:      header.TxHash,
		ReceiptHash: header.ReceiptHash,
		Bloom:       tibcbsc.Bloom(header.Bloom),
		Difficulty:  header.Difficulty,
		Number:      header.Number,
		GasLimit:    header.GasLimit,
		GasUsed:     header.GasUsed,
		Time:        header.Time,
		Extra:       header.Extra,
		MixDigest:   header.MixDigest,
		Nonce:       tibcbsc.BlockNonce(header.Nonce),
	}, nil
}
func GetEthNodeHeader(restClient *RestClient, url string, height uint64) (*tibceth.EthHeader, error) {
	params := []interface{}{fmt.Sprintf("0x%x", height), true}
	req := &blockReq{
		JsonRpc: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  params,
		Id:      1,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetNodeHeight: marshal req err: %s", err)
	}
	rspdata, err := restClient.SendRestRequest(url, reqdata)
	if err != nil {
		return nil, fmt.Errorf("GetNodeHeight err: %s", err)
	}
	rsp := &blockRsp{}
	err = json.Unmarshal(rspdata, rsp)
	if err != nil {
		return nil, fmt.Errorf("GetNodeHeight, unmarshal resp err: %s", err)
	}
	if rsp.Error != nil {
		return nil, fmt.Errorf("GetNodeHeight, return error: %s", rsp.Error.Message)
	}

	if rsp.Result == nil {
		return nil, errors.New("GetNodeHeight, no result")
	}

	header := rsp.Result
	return &tibceth.EthHeader{
		ParentHash:  header.ParentHash,
		UncleHash:   header.UncleHash,
		Coinbase:    header.Coinbase,
		Root:        header.Root,
		TxHash:      header.TxHash,
		ReceiptHash: header.ReceiptHash,
		Bloom:       header.Bloom,
		Difficulty:  header.Difficulty,
		Number:      header.Number,
		GasLimit:    header.GasLimit,
		GasUsed:     header.GasUsed,
		Time:        header.Time,
		Extra:       header.Extra,
		MixDigest:   header.MixDigest,
		Nonce:       header.Nonce,
		BaseFee:     header.BaseFee,
	}, nil
}
