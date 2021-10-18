package tibc_sdk_go

import (
	"context"
	"fmt"

	tibcbsc "github.com/bianjieai/tibc-sdk-go/bsc"
	"github.com/bianjieai/tibc-sdk-go/client"
	commitmenttypes "github.com/bianjieai/tibc-sdk-go/commitment"
	tibceth "github.com/bianjieai/tibc-sdk-go/eth"
	"github.com/bianjieai/tibc-sdk-go/packet"
	"github.com/bianjieai/tibc-sdk-go/tendermint"
	tibcnft "github.com/bianjieai/tibc-sdk-go/types"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	"github.com/irisnet/core-sdk-go/codec"
	cryptotypes "github.com/irisnet/core-sdk-go/codec/types"
	"github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/query"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

type Client struct {
	types.EncodingConfig
	types.BaseClient
}

func NewClient(baseClient types.BaseClient, encodingConfig types.EncodingConfig) Client {
	tibcClient := &Client{
		BaseClient:     baseClient,
		EncodingConfig: encodingConfig,
	}
	tibcClient.RegisterInterfaceTypes(tibcClient.EncodingConfig.InterfaceRegistry)
	return *tibcClient
}

func (c Client) RegisterInterfaceTypes(registry cryptotypes.InterfaceRegistry) {
	packet.RegisterInterfaces(registry)
	tendermint.RegisterInterfaces(registry)
	tibctypes.RegisterInterfaces(registry)
	tibcnft.RegisterInterfaces(registry)
	tibcbsc.RegisterInterfaces(registry)
	tibceth.RegisterInterfaces(registry)
}

func (c Client) Name() string {
	return "tibc"
}

// GetClientState queries an IBC light client.
func (c Client) GetClientState(chainName string) (tibctypes.ClientState, tibctypes.IError) {
	var clientState tibctypes.ClientState
	in := &client.QueryClientStateRequest{
		ChainName: chainName,
	}
	conn, err := c.GenConn()
	if err != nil {
		return clientState, tibctypes.ErrChainConn
	}

	res, err := client.NewQueryClient(conn).ClientState(
		context.Background(),
		in,
	)
	if err != nil {
		return clientState, tibctypes.ErrGetLightClientState
	}

	if err := c.Codec.UnpackAny(res.ClientState, &clientState); err != nil {
		return clientState, tibctypes.ErrUnpackAny
	}
	return clientState, nil

}

// GetClientStates queries all the IBC light clients of a chain.
func (c Client) GetClientStates() ([]tibctypes.ClientState, tibctypes.IError) {
	in := &client.QueryClientStatesRequest{}
	conn, err := c.GenConn()
	if err != nil {
		return nil, tibctypes.ErrChainConn
	}
	res, err := client.NewQueryClient(conn).ClientStates(
		context.Background(),
		in,
	)
	if err != nil {
		return nil, tibctypes.ErrGetLightClientState
	}
	clientState := make([]tibctypes.ClientState, len(res.ClientStates))
	for index, value := range res.ClientStates {
		if err := c.Codec.UnpackAny(value.ClientState, &clientState[index]); err != nil {
			return nil, tibctypes.ErrUnpackAny
		}
	}
	return clientState, nil
}

// GetConsensusState queries a consensus state associated with a client state at
// a given height.
func (c Client) GetConsensusState(chainName string, height uint64) (tibctypes.ConsensusState, tibctypes.IError) {
	req := &client.QueryConsensusStateRequest{
		ChainName:      chainName,
		RevisionHeight: height,
	}
	conn, err := c.GenConn()
	if err != nil {
		return nil, tibctypes.ErrChainConn
	}

	res, err := client.NewQueryClient(conn).ConsensusState(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, tibctypes.ErrGetLightClientConsensusState
	}
	var consensusState tibctypes.ConsensusState

	if err := c.Codec.UnpackAny(res.ConsensusState, &consensusState); err != nil {
		return nil, tibctypes.ErrUnpackAny
	}

	return consensusState, nil
}

// GetConsensusStates queries all the consensus state associated with a given
// client.
func (c Client) GetConsensusStates(chainName string) ([]tibctypes.ConsensusState, tibctypes.IError) {
	req := &client.QueryConsensusStatesRequest{
		ChainName: chainName,
	}
	conn, err := c.GenConn()
	if err != nil {
		return nil, tibctypes.ErrChainConn
	}

	res, err := client.NewQueryClient(conn).ConsensusStates(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, tibctypes.ErrGetLightClientConsensusState
	}
	ConsensusState := make([]tibctypes.ConsensusState, len(res.ConsensusStates))
	for index, value := range res.ConsensusStates {
		if err := c.Codec.UnpackAny(value.ConsensusState, &ConsensusState[index]); err != nil {
			return nil, tibctypes.ErrUnpackAny
		}
	}
	return ConsensusState, nil
}

// Relayers queries all the relayers associated with a given
// client.
func (c Client) Relayers(chainName string) ([]string, tibctypes.IError) {
	req := &client.QueryRelayersRequest{
		ChainName: chainName,
	}
	conn, err := c.GenConn()
	if err != nil {
		return nil, tibctypes.IErrorWrap(tibctypes.ErrChainConn, err.Error())
	}

	relay, err := client.NewQueryClient(conn).Relayers(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, tibctypes.IErrorWrap(tibctypes.ErrGetRelayer, err.Error())
	}

	return relay.Relayers, nil
}

func (c Client) UpdateClient(req tibctypes.UpdateClientRequest, baseTx types.BaseTx) (ctypes.ResultTx, tibctypes.IError) {
	owner, err := c.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {

		return ctypes.ResultTx{}, tibctypes.IErrorWrap(tibctypes.ErrInvalidAddress, err.Error())
	}
	res, errs := cryptotypes.NewAnyWithValue(req.Header)
	if errs != nil {
		return ctypes.ResultTx{}, tibctypes.IErrorWrap(tibctypes.ErrPackAny, err.Error())
	}
	msg := &client.MsgUpdateClient{
		ChainName: req.ChainName,
		// header to update the light client
		Header: res,
		// signer address
		Signer: owner.String(),
	}
	resultTx, err := c.BuildAndSend([]types.Msg{msg}, baseTx)
	if err != nil {
		return ctypes.ResultTx{}, tibctypes.IErrorWrap(tibctypes.ErrUpdateClient, err.Error())
	}
	return resultTx, nil
}

func (c Client) PacketCommitment(destChain string, sourceChain string, sequence uint64) (*packet.QueryPacketCommitmentResponse, tibctypes.IError) {
	req := &packet.QueryPacketCommitmentRequest{
		DestChain:   destChain,
		SourceChain: sourceChain,
		Sequence:    sequence,
	}
	conn, err := c.GenConn()
	if err != nil {
		return nil, tibctypes.IErrorWrap(tibctypes.ErrChainConn, err.Error())
	}
	req1, err := packet.NewQueryClient(conn).PacketCommitment(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, tibctypes.IErrorWrap(tibctypes.ErrGetCommitmentPacket, err.Error())
	}
	return req1, nil
}

func (c Client) PacketCommitments(destChain string, sourceChain string, Pagination *query.PageRequest) (*packet.QueryPacketCommitmentsResponse, tibctypes.IError) {
	req := &packet.QueryPacketCommitmentsRequest{
		DestChain:   destChain,
		SourceChain: sourceChain,
		Pagination:  Pagination,
	}
	conn, err := c.GenConn()
	if err != nil {
		return nil, tibctypes.IErrorWrap(tibctypes.ErrChainConn, err.Error())
	}
	packComms, err := packet.NewQueryClient(conn).PacketCommitments(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, tibctypes.IErrorWrap(tibctypes.ErrGetCommitmentPacket, err.Error())
	}
	return packComms, nil
}

func (c Client) PacketReceipt(destChain string, sourceChain string, sequence uint64) (*packet.QueryPacketReceiptResponse, tibctypes.IError) {
	req := &packet.QueryPacketReceiptRequest{
		DestChain:   destChain,
		SourceChain: sourceChain,
		Sequence:    sequence,
	}
	conn, err := c.GenConn()
	if err != nil {
		return nil, tibctypes.IErrorWrap(tibctypes.ErrChainConn, err.Error())
	}
	receipt, err := packet.NewQueryClient(conn).PacketReceipt(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, tibctypes.IErrorWrap(tibctypes.ErrGetReceiptPacket, err.Error())
	}
	return receipt, nil
}
func (c Client) PacketAcknowledgement(destChain string, sourceChain string, sequence uint64) (*packet.QueryPacketAcknowledgementResponse, tibctypes.IError) {
	req := &packet.QueryPacketAcknowledgementRequest{
		DestChain:   destChain,
		SourceChain: sourceChain,
		Sequence:    sequence,
	}
	conn, err := c.GenConn()
	if err != nil {
		return nil, tibctypes.IErrorWrap(tibctypes.ErrChainConn, err.Error())
	}
	acknowledgement, err := packet.NewQueryClient(conn).PacketAcknowledgement(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, tibctypes.IErrorWrap(tibctypes.ErrGetAckPacket, err.Error())
	}
	return acknowledgement, nil
}
func (c Client) PacketAcknowledgements(destChain string, sourceChain string, Pagination *query.PageRequest) (*packet.QueryPacketAcknowledgementsResponse, tibctypes.IError) {
	req := &packet.QueryPacketAcknowledgementsRequest{
		DestChain:   destChain,
		SourceChain: sourceChain,
		Pagination:  Pagination,
	}
	conn, err := c.GenConn()
	if err != nil {
		return nil, tibctypes.IErrorWrap(tibctypes.ErrChainConn, err.Error())
	}
	acknowledgements, err := packet.NewQueryClient(conn).PacketAcknowledgements(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, tibctypes.IErrorWrap(tibctypes.ErrGetAckPacket, err.Error())
	}
	return acknowledgements, nil
}
func (c Client) UnreceivedPackets(destChain string, sourceChain string, packetCommitmentSequences []uint64) (*packet.QueryUnreceivedPacketsResponse, tibctypes.IError) {
	req := &packet.QueryUnreceivedPacketsRequest{
		DestChain:                 destChain,
		SourceChain:               sourceChain,
		PacketCommitmentSequences: packetCommitmentSequences,
	}
	conn, err := c.GenConn()
	if err != nil {
		return nil, tibctypes.IErrorWrap(tibctypes.ErrChainConn, err.Error())
	}
	unreceivedPackets, err := packet.NewQueryClient(conn).UnreceivedPackets(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, tibctypes.IErrorWrap(tibctypes.ErrGetUnreceivedPacket, err.Error())
	}
	return unreceivedPackets, nil
}

func (c Client) UnreceivedAcks(destChain string, sourceChain string, packetAckSequences []uint64) (*packet.QueryUnreceivedAcksResponse, tibctypes.IError) {
	req := &packet.QueryUnreceivedAcksRequest{
		DestChain:          destChain,
		SourceChain:        sourceChain,
		PacketAckSequences: packetAckSequences,
	}
	conn, err := c.GenConn()
	if err != nil {
		return nil, tibctypes.IErrorWrap(tibctypes.ErrChainConn, err.Error())
	}
	unreceivedAcks, err := packet.NewQueryClient(conn).UnreceivedAcks(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, tibctypes.IErrorWrap(tibctypes.ErrGetUnreceivedPacket, err.Error())
	}
	return unreceivedAcks, nil
}
func (c Client) RecvPackets(msgs []types.Msg, baseTx types.BaseTx) (ctypes.ResultTx, tibctypes.IError) {
	txreq, err := c.BuildAndSend(msgs, baseTx)
	if err != nil {
		return ctypes.ResultTx{}, tibctypes.IErrorWrap(tibctypes.ErrRecvPacket, err.Error())
	}
	return txreq, nil
}

func (c Client) RecvPacket(proof []byte, pack packet.Packet, height int64, revisionNumber uint64, baseTx types.BaseTx) (ctypes.ResultTx, tibctypes.IError) {
	owner, err := c.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, tibctypes.IErrorWrap(tibctypes.ErrInvalidAddress, err.Error())
	}
	msg := &packet.MsgRecvPacket{
		Packet:          pack,
		ProofCommitment: proof,
		ProofHeight: client.Height{
			RevisionNumber: revisionNumber,
			RevisionHeight: uint64(height),
		},
		Signer: owner.String(),
	}
	txreq, err := c.BuildAndSend([]types.Msg{msg}, baseTx)
	if err != nil {
		return ctypes.ResultTx{}, tibctypes.IErrorWrap(tibctypes.ErrRecvPacket, err.Error())
	}
	return txreq, nil
}

func (c Client) Acknowledgement(proof []byte, acknowledgement []byte, pack packet.Packet, height int64, revisionNumber uint64, baseTx types.BaseTx) (ctypes.ResultTx, tibctypes.IError) {
	owner, err := c.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, tibctypes.IErrorWrap(tibctypes.ErrInvalidAddress, err.Error())
	}
	msg := &packet.MsgAcknowledgement{
		Packet:          pack,
		Acknowledgement: acknowledgement,
		ProofAcked:      proof,
		ProofHeight: client.Height{
			RevisionNumber: revisionNumber,
			RevisionHeight: uint64(height),
		},
		Signer: owner.String(),
	}
	txreq, err := c.BuildAndSend([]types.Msg{msg}, baseTx)
	if err != nil {
		return ctypes.ResultTx{}, tibctypes.IErrorWrap(tibctypes.ErrSendAckPacket, err.Error())
	}
	return txreq, nil
}

func (c Client) CleanPacket(cleanPacket packet.CleanPacket, baseTx types.BaseTx) (ctypes.ResultTx, tibctypes.IError) {
	owner, err := c.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, tibctypes.IErrorWrap(tibctypes.ErrInvalidAddress, err.Error())
	}
	msg := &packet.MsgCleanPacket{
		CleanPacket: cleanPacket,
		Signer:      owner.String(),
	}
	txreq, err := c.BuildAndSend([]types.Msg{msg}, baseTx)
	if err != nil {
		return ctypes.ResultTx{}, tibctypes.IErrorWrap(tibctypes.ErrSendCleanPacket, err.Error())
	}
	return txreq, nil
}

func (c Client) RecvCleanPacket(proof []byte, pack packet.CleanPacket, height int64, revisionNumber uint64, baseTx types.BaseTx) (ctypes.ResultTx, tibctypes.IError) {
	owner, err := c.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, tibctypes.IErrorWrap(tibctypes.ErrInvalidAddress, err.Error())
	}

	msg := &packet.MsgRecvCleanPacket{
		CleanPacket:     pack,
		ProofCommitment: proof,
		ProofHeight: client.Height{
			RevisionNumber: revisionNumber,
			RevisionHeight: uint64(height),
		},
		Signer: owner.String(),
	}
	txreq, err := c.BuildAndSend([]types.Msg{msg}, baseTx)
	if err != nil {
		return ctypes.ResultTx{}, tibctypes.IErrorWrap(tibctypes.ErrRecvCleanPacket, err.Error())
	}
	return txreq, nil
}

func (c Client) NftTransfer(class, id, receiver, destChainName, realayChainName string, baseTx types.BaseTx) (ctypes.ResultTx, tibctypes.IError) {
	owner, err := c.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, tibctypes.IErrorWrap(tibctypes.ErrInvalidAddress, err.Error())
	}
	msg := &tibcnft.MsgNftTransfer{
		Class:       class,
		Id:          id,
		Sender:      owner.String(),
		Receiver:    receiver,
		DestChain:   destChainName,
		RealayChain: realayChainName,
	}
	txreq, err := c.BuildAndSend([]types.Msg{msg}, baseTx)
	if err != nil {
		return ctypes.ResultTx{}, tibctypes.IErrorWrap(tibctypes.ErrNftTransfer, err.Error())
	}
	return txreq, nil
}
func (c Client) QueryTendermintProof(height int64, key []byte) ([]byte, []byte, uint64, error) {
	// ABCI queries at heights 1, 2 or less than or equal to 0 are not supported.
	// Base app does not support queries for height less than or equal to 1.
	// Therefore, a query at height 2 would be equivalent to a query at height 3.
	// A height of 0 will query with the lastest state.
	if height != 0 && height <= 2 {
		return nil, nil, 0, fmt.Errorf("proof queries at height <= 2 are not supported")
	}
	// Use the IAVL height if a valid tendermint height is passed in.
	// A height of 0 will query with the latest state.
	if height != 0 {
		height--
	}
	res, err := c.QueryStore(key, "tibc", height, true)
	if err != nil {
		return nil, nil, 0, err
	}

	merkleProof, err := commitmenttypes.ConvertProofs(res.ProofOps)
	if err != nil {
		return nil, nil, 0, err
	}
	cdc := codec.NewProtoCodec(c.EncodingConfig.InterfaceRegistry)

	proofBz, err := cdc.Marshal(&merkleProof)
	if err != nil {
		return nil, nil, 0, err
	}

	return res.Value, proofBz, uint64(res.Height) + 1, nil
}
