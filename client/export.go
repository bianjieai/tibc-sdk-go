package client

import (
	"github.com/bianjieai/tibc-sdk-go/packet"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	sdk "github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/query"
)

type ChainClient interface {
	sdk.Module

	// lightClient
	GetClientState(chainName string) (tibctypes.ClientState, error)
	GetClientStates() ([]tibctypes.ClientState, error)
	GetConsensusState(chainName string, height uint64) (tibctypes.ConsensusState, error)
	GetConsensusStates(chainName string) ([]tibctypes.ConsensusState, error)
	Relayers(chainName string) ([]string, error)
	UpdateClient(tibctypes.UpdateClientRequest, sdk.BaseTx) (sdk.ResultTx, sdk.Error)

	//  packet
	PacketCommitment(destChain string, sourceChain string, sequence uint64) (*packet.QueryPacketCommitmentResponse, error)
	PacketCommitments(destChain string, sourceChain string, Pagination *query.PageRequest) (*packet.QueryPacketCommitmentsResponse, error)
	PacketReceipt(destChain string, sourceChain string, sequence uint64) (*packet.QueryPacketReceiptResponse, error)
	PacketAcknowledgement(destChain string, sourceChain string, sequence uint64) (*packet.QueryPacketAcknowledgementResponse, error)
	PacketAcknowledgements(destChain string, sourceChain string, Pagination *query.PageRequest) (*packet.QueryPacketAcknowledgementsResponse, error)
	UnreceivedPackets(destChain string, sourceChain string, packetCommitmentSequences []uint64) (*packet.QueryUnreceivedPacketsResponse, error)
	UnreceivedAcks(destChain string, sourceChain string, packetAckSequences []uint64) (*packet.QueryUnreceivedAcksResponse, error)
	NextSequenceReceive(destChain string, sourceChain string) (*packet.QueryNextSequenceReceiveResponse, error)
}
