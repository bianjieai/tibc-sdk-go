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
func (c Client) GetClientState(chainName string) (tibctypes.ClientState, error) {
	var clientState tibctypes.ClientState
	in := &client.QueryClientStateRequest{
		ChainName: chainName,
	}

	conn, err := c.CoreSdk.GenConn()
	if err != nil {
		return clientState, types.Wrap(err)
	}

	res, err := client.NewQueryClient(conn).ClientState(
		context.Background(),
		in,
	)
	if err != nil {
		return clientState, types.Wrap(err)
	}

	if err := c.Marshaler.UnpackAny(res.ClientState, &clientState); err != nil {
		return clientState, types.Wrap(err)
	}
	return clientState, nil

}

// GetClientStates queries all the IBC light clients of a chain.
func (c Client) GetClientStates() ([]tibctypes.ClientState, error) {
	in := &client.QueryClientStatesRequest{}
	conn, err := c.CoreSdk.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}
	res, err := client.NewQueryClient(conn).ClientStates(
		context.Background(),
		in,
	)
	if err != nil {
		return nil, types.Wrap(err)
	}
	clientState := make([]tibctypes.ClientState, len(res.ClientStates))
	for index, value := range res.ClientStates {
		if err := c.Marshaler.UnpackAny(value.ClientState, &clientState[index]); err != nil {
			return nil, types.Wrap(err)
		}
	}

	return clientState, err
}

// GetConsensusState queries a consensus state associated with a client state at
// a given height.
func (c Client) GetConsensusState(chainName string, height uint64) (tibctypes.ConsensusState, error) {
	req := &client.QueryConsensusStateRequest{
		ChainName:      chainName,
		RevisionHeight: height,
	}

	conn, err := c.CoreSdk.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}

	res, err := client.NewQueryClient(conn).ConsensusState(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, types.Wrap(err)
	}
	var consensusState tibctypes.ConsensusState

	if err := c.Marshaler.UnpackAny(res.ConsensusState, &consensusState); err != nil {
		return nil, types.Wrap(err)
	}

	return consensusState, nil
}

// GetConsensusStates queries all the consensus state associated with a given
// client.
func (c Client) GetConsensusStates(chainName string) ([]tibctypes.ConsensusState, error) {
	req := &client.QueryConsensusStatesRequest{
		ChainName: chainName,
	}
	conn, err := c.CoreSdk.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}

	res, err := client.NewQueryClient(conn).ConsensusStates(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, types.Wrap(err)
	}
	ConsensusState := make([]tibctypes.ConsensusState, len(res.ConsensusStates))
	for index, value := range res.ConsensusStates {
		if err := c.Marshaler.UnpackAny(value.ConsensusState, &ConsensusState[index]); err != nil {
			return nil, types.Wrap(err)
		}
	}
	return ConsensusState, nil
}

// Relayers queries all the relayers associated with a given
// client.
func (c Client) Relayers(chainName string) ([]string, error) {
	req := &client.QueryRelayersRequest{
		ChainName: chainName,
	}
	conn, err := c.CoreSdk.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}

	relay, err := client.NewQueryClient(conn).Relayers(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, types.Wrap(err)
	}

	return relay.Relayers, nil
}

func (c Client) UpdateClient(req tibctypes.UpdateClientRequest, baseTx types.BaseTx) (types.ResultTx, types.Error) {
	owner, err := c.CoreSdk.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return types.ResultTx{}, types.Wrap(err)
	}
	res, errs := cryptotypes.NewAnyWithValue(req.Header)
	if errs != nil {
		return types.ResultTx{}, types.Wrap(errs)
	}
	msg := &client.MsgUpdateClient{
		ChainName: req.ChainName,
		// header to update the light client
		Header: res,
		// signer address
		Signer: owner.String(),
	}

	return c.CoreSdk.BuildAndSend([]types.Msg{msg}, baseTx)
}

func (c Client) PacketCommitment(destChain string, sourceChain string, sequence uint64) (*packet.QueryPacketCommitmentResponse, error) {
	req := &packet.QueryPacketCommitmentRequest{
		DestChain:   destChain,
		SourceChain: sourceChain,
		Sequence:    sequence,
	}
	conn, err := c.CoreSdk.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}
	return packet.NewQueryClient(conn).PacketCommitment(
		context.Background(),
		req,
	)
}

func (c Client) PacketCommitments(destChain string, sourceChain string, Pagination *query.PageRequest) (*packet.QueryPacketCommitmentsResponse, error) {
	req := &packet.QueryPacketCommitmentsRequest{
		DestChain:   destChain,
		SourceChain: sourceChain,
		Pagination:  Pagination,
	}
	conn, err := c.CoreSdk.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}
	return packet.NewQueryClient(conn).PacketCommitments(
		context.Background(),
		req,
	)
}

func (c Client) PacketReceipt(destChain string, sourceChain string, sequence uint64) (*packet.QueryPacketReceiptResponse, error) {
	req := &packet.QueryPacketReceiptRequest{
		DestChain:   destChain,
		SourceChain: sourceChain,
		Sequence:    sequence,
	}
	conn, err := c.CoreSdk.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}
	return packet.NewQueryClient(conn).PacketReceipt(
		context.Background(),
		req,
	)
}
func (c Client) PacketAcknowledgement(destChain string, sourceChain string, sequence uint64) (*packet.QueryPacketAcknowledgementResponse, error) {
	req := &packet.QueryPacketAcknowledgementRequest{
		DestChain:   destChain,
		SourceChain: sourceChain,
		Sequence:    sequence,
	}
	conn, err := c.CoreSdk.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}
	return packet.NewQueryClient(conn).PacketAcknowledgement(
		context.Background(),
		req,
	)
}
func (c Client) PacketAcknowledgements(destChain string, sourceChain string, Pagination *query.PageRequest) (*packet.QueryPacketAcknowledgementsResponse, error) {
	req := &packet.QueryPacketAcknowledgementsRequest{
		DestChain:   destChain,
		SourceChain: sourceChain,
		Pagination:  Pagination,
	}
	conn, err := c.CoreSdk.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}
	return packet.NewQueryClient(conn).PacketAcknowledgements(
		context.Background(),
		req,
	)
}
func (c Client) UnreceivedPackets(destChain string, sourceChain string, packetCommitmentSequences []uint64) (*packet.QueryUnreceivedPacketsResponse, error) {
	req := &packet.QueryUnreceivedPacketsRequest{
		DestChain:                 destChain,
		SourceChain:               sourceChain,
		PacketCommitmentSequences: packetCommitmentSequences,
	}
	conn, err := c.CoreSdk.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}
	return packet.NewQueryClient(conn).UnreceivedPackets(
		context.Background(),
		req,
	)
}

func (c Client) UnreceivedAcks(destChain string, sourceChain string, packetAckSequences []uint64) (*packet.QueryUnreceivedAcksResponse, error) {
	req := &packet.QueryUnreceivedAcksRequest{
		DestChain:          destChain,
		SourceChain:        sourceChain,
		PacketAckSequences: packetAckSequences,
	}
	conn, err := c.CoreSdk.GenConn()
	if err != nil {
		return nil, types.Wrap(err)
	}
	return packet.NewQueryClient(conn).UnreceivedAcks(
		context.Background(),
		req,
	)
}
func (c Client) RecvPackets(msgs []types.Msg, baseTx types.BaseTx) (types.ResultTx, types.Error) {
	return c.CoreSdk.BuildAndSend(msgs, baseTx)
}

func (c Client) RecvPacket(proof []byte, pack packet.Packet, height int64, revisionNumber uint64, baseTx types.BaseTx) (types.ResultTx, types.Error) {
	owner, err := c.CoreSdk.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return types.ResultTx{}, types.Wrap(err)
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
	return c.CoreSdk.BuildAndSend([]types.Msg{msg}, baseTx)
}

func (c Client) Acknowledgement(proof []byte, acknowledgement []byte, pack packet.Packet, height int64, revisionNumber uint64, baseTx types.BaseTx) (types.ResultTx, types.Error) {
	owner, err := c.CoreSdk.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return types.ResultTx{}, types.Wrap(err)
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
	return c.CoreSdk.BuildAndSend([]types.Msg{msg}, baseTx)
}

func (c Client) CleanPacket(cleanPacket packet.CleanPacket, baseTx types.BaseTx) (types.ResultTx, types.Error) {
	owner, err := c.CoreSdk.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return types.ResultTx{}, types.Wrap(err)
	}
	msg := &packet.MsgCleanPacket{
		CleanPacket: cleanPacket,
		Signer:      owner.String(),
	}
	return c.CoreSdk.BuildAndSend([]types.Msg{msg}, baseTx)
}

func (c Client) RecvCleanPacket(proof []byte, pack packet.CleanPacket, height int64, revisionNumber uint64, baseTx types.BaseTx) (types.ResultTx, types.Error) {
	owner, err := c.CoreSdk.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return types.ResultTx{}, types.Wrap(err)
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
	return c.CoreSdk.BuildAndSend([]types.Msg{msg}, baseTx)
}
