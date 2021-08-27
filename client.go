package tibc_sdk_go

import (
	"context"

	"github.com/bianjieai/tibc-sdk-go/client"
	"github.com/bianjieai/tibc-sdk-go/packet"
	"github.com/bianjieai/tibc-sdk-go/tendermint"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	sdk "github.com/irisnet/core-sdk-go"
	commoncodec "github.com/irisnet/core-sdk-go/common/codec"
	cryptotypes "github.com/irisnet/core-sdk-go/common/codec/types"
	"github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/query"
)

type Client struct {
	CoreSdk sdk.Client
	commoncodec.Marshaler
}

func NewClient(coreClient sdk.Client) Client {
	tibcClient := &Client{
		CoreSdk:   coreClient,
		Marshaler: coreClient.EncodingConfig().Marshaler,
	}
	tibcClient.RegisterInterfaceTypes(coreClient.EncodingConfig().InterfaceRegistry)
	return *tibcClient
}

func (c Client) RegisterInterfaceTypes(registry cryptotypes.InterfaceRegistry) {
	packet.RegisterInterfaces(registry)
	tendermint.RegisterInterfaces(registry)
	tibctypes.RegisterInterfaces(registry)
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

	conn, err := c.CoreSdk.GenConn()
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

	if err := c.Marshaler.UnpackAny(res.ClientState, &clientState); err != nil {
		return clientState, tibctypes.ErrUnpackAny
	}
	return clientState, nil

}

// GetClientStates queries all the IBC light clients of a chain.
func (c Client) GetClientStates() ([]tibctypes.ClientState, tibctypes.IError) {
	in := &client.QueryClientStatesRequest{}
	conn, err := c.CoreSdk.GenConn()
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
		if err := c.Marshaler.UnpackAny(value.ClientState, &clientState[index]); err != nil {
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

	conn, err := c.CoreSdk.GenConn()
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

	if err := c.Marshaler.UnpackAny(res.ConsensusState, &consensusState); err != nil {
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
	conn, err := c.CoreSdk.GenConn()
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
		if err := c.Marshaler.UnpackAny(value.ConsensusState, &ConsensusState[index]); err != nil {
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
	conn, err := c.CoreSdk.GenConn()
	if err != nil {
		return nil, tibctypes.ErrChainConn
	}

	relay, err := client.NewQueryClient(conn).Relayers(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, tibctypes.ErrGetRelayer
	}

	return relay.Relayers, nil
}

func (c Client) UpdateClient(req tibctypes.UpdateClientRequest, baseTx types.BaseTx) (types.ResultTx, tibctypes.IError) {
	owner, err := c.CoreSdk.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return types.ResultTx{}, err.(tibctypes.IError)
	}
	res, errs := cryptotypes.NewAnyWithValue(req.Header)
	if errs != nil {
		return types.ResultTx{}, tibctypes.ErrPackAny
	}
	msg := &client.MsgUpdateClient{
		ChainName: req.ChainName,
		// header to update the light client
		Header: res,
		// signer address
		Signer: owner.String(),
	}
	resultTx, err := c.CoreSdk.BuildAndSend([]types.Msg{msg}, baseTx)
	if err != nil {
		return types.ResultTx{}, tibctypes.ErrUpdateClient
	}
	return resultTx, nil
}

func (c Client) PacketCommitment(destChain string, sourceChain string, sequence uint64) (*packet.QueryPacketCommitmentResponse, tibctypes.IError) {
	req := &packet.QueryPacketCommitmentRequest{
		DestChain:   destChain,
		SourceChain: sourceChain,
		Sequence:    sequence,
	}
	conn, err := c.CoreSdk.GenConn()
	if err != nil {
		return nil, tibctypes.ErrChainConn
	}
	req1, err := packet.NewQueryClient(conn).PacketCommitment(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, tibctypes.ErrGetCommitmentPacket
	}
	return req1, nil
	//return req1, tibctypes.ErrGetCommitmentPacket
}

func (c Client) PacketCommitments(destChain string, sourceChain string, Pagination *query.PageRequest) (*packet.QueryPacketCommitmentsResponse, tibctypes.IError) {
	req := &packet.QueryPacketCommitmentsRequest{
		DestChain:   destChain,
		SourceChain: sourceChain,
		Pagination:  Pagination,
	}
	conn, err := c.CoreSdk.GenConn()
	if err != nil {
		return nil, tibctypes.ErrChainConn
	}
	packComms, err := packet.NewQueryClient(conn).PacketCommitments(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, tibctypes.ErrGetCommitmentPacket
	}
	return packComms, nil
}

func (c Client) PacketReceipt(destChain string, sourceChain string, sequence uint64) (*packet.QueryPacketReceiptResponse, tibctypes.IError) {
	req := &packet.QueryPacketReceiptRequest{
		DestChain:   destChain,
		SourceChain: sourceChain,
		Sequence:    sequence,
	}
	conn, err := c.CoreSdk.GenConn()
	if err != nil {
		return nil, tibctypes.ErrChainConn
	}
	receipt, err := packet.NewQueryClient(conn).PacketReceipt(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, tibctypes.ErrGetReceiptPacket
	}
	return receipt, nil
}
func (c Client) PacketAcknowledgement(destChain string, sourceChain string, sequence uint64) (*packet.QueryPacketAcknowledgementResponse, tibctypes.IError) {
	req := &packet.QueryPacketAcknowledgementRequest{
		DestChain:   destChain,
		SourceChain: sourceChain,
		Sequence:    sequence,
	}
	conn, err := c.CoreSdk.GenConn()
	if err != nil {
		return nil, tibctypes.ErrChainConn
	}
	acknowledgement, err := packet.NewQueryClient(conn).PacketAcknowledgement(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, tibctypes.ErrGetAckPacket
	}
	return acknowledgement, nil
}
func (c Client) PacketAcknowledgements(destChain string, sourceChain string, Pagination *query.PageRequest) (*packet.QueryPacketAcknowledgementsResponse, tibctypes.IError) {
	req := &packet.QueryPacketAcknowledgementsRequest{
		DestChain:   destChain,
		SourceChain: sourceChain,
		Pagination:  Pagination,
	}
	conn, err := c.CoreSdk.GenConn()
	if err != nil {
		return nil, tibctypes.ErrChainConn
	}
	acknowledgements, err := packet.NewQueryClient(conn).PacketAcknowledgements(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, tibctypes.ErrGetAckPacket
	}
	return acknowledgements, nil
}
func (c Client) UnreceivedPackets(destChain string, sourceChain string, packetCommitmentSequences []uint64) (*packet.QueryUnreceivedPacketsResponse, tibctypes.IError) {
	req := &packet.QueryUnreceivedPacketsRequest{
		DestChain:                 destChain,
		SourceChain:               sourceChain,
		PacketCommitmentSequences: packetCommitmentSequences,
	}
	conn, err := c.CoreSdk.GenConn()
	if err != nil {
		return nil, tibctypes.ErrChainConn
	}
	unreceivedPackets, err := packet.NewQueryClient(conn).UnreceivedPackets(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, tibctypes.ErrGetUnreceivedPacket
	}
	return unreceivedPackets, nil
}

func (c Client) UnreceivedAcks(destChain string, sourceChain string, packetAckSequences []uint64) (*packet.QueryUnreceivedAcksResponse, tibctypes.IError) {
	req := &packet.QueryUnreceivedAcksRequest{
		DestChain:          destChain,
		SourceChain:        sourceChain,
		PacketAckSequences: packetAckSequences,
	}
	conn, err := c.CoreSdk.GenConn()
	if err != nil {
		return nil, tibctypes.ErrChainConn
	}
	unreceivedAcks, err := packet.NewQueryClient(conn).UnreceivedAcks(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, tibctypes.ErrGetUnreceivedPacket
	}
	return unreceivedAcks, nil
}
func (c Client) RecvPackets(msgs []types.Msg, baseTx types.BaseTx) (types.ResultTx, tibctypes.IError) {

	txreq, err := c.CoreSdk.BuildAndSend(msgs, baseTx)
	if err != nil {
		return types.ResultTx{}, tibctypes.ErrRecvPacket
	}
	return txreq, nil
}

func (c Client) RecvPacket(proof []byte, pack packet.Packet, height int64, revisionNumber uint64, baseTx types.BaseTx) (types.ResultTx, tibctypes.IError) {
	owner, err := c.CoreSdk.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return types.ResultTx{}, err.(tibctypes.IError)
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
	txreq, err := c.CoreSdk.BuildAndSend([]types.Msg{msg}, baseTx)
	if err != nil {
		return types.ResultTx{}, tibctypes.ErrRecvPacket
	}
	return txreq, nil
}

func (c Client) Acknowledgement(proof []byte, acknowledgement []byte, pack packet.Packet, height int64, revisionNumber uint64, baseTx types.BaseTx) (types.ResultTx, tibctypes.IError) {
	owner, err := c.CoreSdk.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return types.ResultTx{}, err.(tibctypes.IError)
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
	txreq, err := c.CoreSdk.BuildAndSend([]types.Msg{msg}, baseTx)
	if err != nil {
		return types.ResultTx{}, tibctypes.ErrSendAckPacket
	}
	return txreq, nil
}

func (c Client) CleanPacket(cleanPacket packet.CleanPacket, baseTx types.BaseTx) (types.ResultTx, tibctypes.IError) {
	owner, err := c.CoreSdk.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return types.ResultTx{}, err.(tibctypes.IError)
	}
	msg := &packet.MsgCleanPacket{
		CleanPacket: cleanPacket,
		Signer:      owner.String(),
	}
	txreq, err := c.CoreSdk.BuildAndSend([]types.Msg{msg}, baseTx)
	if err != nil {
		return types.ResultTx{}, tibctypes.ErrSendCleanPacket
	}
	return txreq, nil
}

func (c Client) RecvCleanPacket(proof []byte, pack packet.CleanPacket, height int64, revisionNumber uint64, baseTx types.BaseTx) (types.ResultTx, tibctypes.IError) {
	owner, err := c.CoreSdk.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return types.ResultTx{}, err.(tibctypes.IError)
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
	txreq, err := c.CoreSdk.BuildAndSend([]types.Msg{msg}, baseTx)
	if err != nil {
		return types.ResultTx{}, tibctypes.ErrRecvCleanPacket
	}
	return txreq, nil
}
