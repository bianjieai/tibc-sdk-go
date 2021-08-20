package integration

import (
	"encoding/hex"
	"fmt"
	"strconv"

	tibc "github.com/bianjieai/tibc-sdk-go"
	"github.com/bianjieai/tibc-sdk-go/packet"
	"github.com/bianjieai/tibc-sdk-go/tendermint"
	"github.com/irisnet/core-sdk-go/types"
)

func queryCommitment(client tibc.Client) *packet.QueryPacketCommitmentResponse {
	res, err := client.PacketCommitment("testCreateClientC", "testCreateClientA", 1)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(res.String())

	res1, err := client.PacketCommitments("testCreateClientC", "testCreateClientA", nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("commitments:  ", res1.Commitments)
	return res
}

func packetReceipt(client tibc.Client) {
	res, err := client.PacketReceipt("testCreateClientC", "testCreateClientA", 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Received)
}
func queryack(client tibc.Client) {
	res, err := client.PacketAcknowledgement("testCreateClientC", "testCreateClientA", 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	res1, err := client.PacketAcknowledgements("testCreateClientC", "testCreateClientA", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.String())
	fmt.Println(res1.String())
}
func queryUnreceivedPacketsAndAcks(client tibc.Client) {
	res, err := client.UnreceivedPackets("testCreateClientA", "testCreateClientC", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("packets :", res.String())
	res1, err := client.UnreceivedAcks("testCreateClientA", "testCreateClientC", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("acks ", res1.String())

}
func getpacket(tx types.ResultQueryTx) (packet.Packet, error) {
	sequence, err := tx.Result.Events.GetValue("send_packet", "packet_sequence")
	if err != nil {
		fmt.Println(err)
		return packet.Packet{}, err
	}
	sourceChain, err := tx.Result.Events.GetValue("send_packet", "packet_src_chain")
	if err != nil {
		fmt.Println(err)
		return packet.Packet{}, err
	}
	destinationChain, err := tx.Result.Events.GetValue("send_packet", "packet_dst_port")
	if err != nil {
		fmt.Println(err)
		return packet.Packet{}, err
	}
	port, err := tx.Result.Events.GetValue("send_packet", "packet_port")
	if err != nil {
		fmt.Println(err)
		return packet.Packet{}, err
	}
	relayChain, err := tx.Result.Events.GetValue("send_packet", "packet_relay_channel")
	if err != nil {
		fmt.Println(err)
		return packet.Packet{}, err
	}
	data, err := tx.Result.Events.GetValue("send_packet", "packet_data")
	if err != nil {
		fmt.Println(err)
		return packet.Packet{}, err
	}
	num, err := strconv.Atoi(sequence)
	if err != nil {
		fmt.Println(err)
		return packet.Packet{}, err
	}
	fmt.Println(num)
	return packet.Packet{
		Sequence:         uint64(num),
		SourceChain:      sourceChain,
		DestinationChain: destinationChain,
		Port:             port,
		RelayChain:       relayChain,
		Data:             []byte(data),
	}, nil
}
func getpacketAndAck(tx types.ResultQueryTx) (packet.Packet, []byte, error) {
	sequence, err := tx.Result.Events.GetValue("write_acknowledgement", "packet_sequence")
	if err != nil {
		fmt.Println(err)
		return packet.Packet{}, nil, err
	}
	sourceChain, err := tx.Result.Events.GetValue("write_acknowledgement", "packet_src_chain")
	if err != nil {
		fmt.Println(err)
		return packet.Packet{}, nil, err
	}
	destinationChain, err := tx.Result.Events.GetValue("write_acknowledgement", "packet_dst_port")
	if err != nil {
		fmt.Println(err)
		return packet.Packet{}, nil, err
	}
	port, err := tx.Result.Events.GetValue("write_acknowledgement", "packet_port")
	if err != nil {
		fmt.Println(err)
		return packet.Packet{}, nil, err
	}
	relayChain, err := tx.Result.Events.GetValue("write_acknowledgement", "packet_relay_channel")
	if err != nil {
		fmt.Println(err)
		return packet.Packet{}, nil, err
	}
	data, err := tx.Result.Events.GetValue("write_acknowledgement", "packet_data")
	if err != nil {
		fmt.Println(err)
		return packet.Packet{}, nil, err
	}
	ack, err := tx.Result.Events.GetValue("write_acknowledgement", "packet_ack")
	if err != nil {
		fmt.Println(err)
		return packet.Packet{}, nil, err
	}
	num, err := strconv.Atoi(sequence)
	if err != nil {
		fmt.Println(err)
		return packet.Packet{}, nil, err
	}
	fmt.Println(num)
	return packet.Packet{
		Sequence:         uint64(num),
		SourceChain:      sourceChain,
		DestinationChain: destinationChain,
		Port:             port,
		RelayChain:       relayChain,
		Data:             []byte(data),
	}, []byte(ack), nil
}
func sendAck(sourceClient tibc.Client, destClient tibc.Client, keyname string) {
	tx, err := sourceClient.CoreSdk.QueryTx("CC57E9C818F1C164CF377674642A4CA67B4F49CDD3648F1C764E7729419C8E82")
	if err != nil {
		fmt.Println(err)
		return
	}
	clients, err := destClient.GetClientState("testCreateClientC")
	height := clients.GetLatestHeight()
	packet1, ack, err := getpacketAndAck(tx)
	ack1, _ := sourceClient.PacketAcknowledgement("testCreateClientC", "testCreateClientA", 1)
	fmt.Println("ack1: ", ack1.String())
	baseTx := types.BaseTx{
		From:               keyname,
		Gas:                0,
		Memo:               "TEST",
		Mode:               types.Commit,
		Password:           "12345678",
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}
	// ProofCommitment and ProofHeight are derived from the packet
	key := packet.PacketAcknowledgementKey(packet1.GetSourceChain(), packet1.GetDestChain(), packet1.GetSequence())
	v1, proofBz, height1, err1 := tendermint.QueryTendermintProof(sourceClient.CoreSdk, int64(height.GetRevisionHeight()), key)

	fmt.Println("ack", hex.EncodeToString(ack))
	fmt.Println("v1", hex.EncodeToString(v1))

	fmt.Println("height1", height1)
	fmt.Println("height", height.GetRevisionHeight())

	fmt.Println("proof ::!!!", proofBz)
	fmt.Println("height::!! ", height.GetRevisionHeight())
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	ress, err := destClient.Acknowledgement(proofBz, ack, packet1, int64(height.GetRevisionHeight()), 0, baseTx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ress)
}

func packetRecive(sourceClient tibc.Client, destClient tibc.Client, keyname string) {
	tx, err := sourceClient.CoreSdk.QueryTx("627D11B87B651F16EBF95C485E561E89A0E299AA4A773E95AB29CEE23F6F08A1")
	if err != nil {
		fmt.Println(err)
		return
	}
	res := queryCommitment(sourceClient)
	fmt.Println("proof : ", res.Proof)
	clients, err := destClient.GetClientState("testCreateClientA")
	height := clients.GetLatestHeight()
	packet1, err := getpacket(tx)
	baseTx := types.BaseTx{
		From:               keyname,
		Gas:                0,
		Memo:               "TEST",
		Mode:               types.Commit,
		Password:           "12345678",
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}

	// ProofCommitment and ProofHeight are derived from the packet
	key := packet.PacketCommitmentKey(packet1.GetSourceChain(), packet1.GetDestChain(), packet1.GetSequence())
	v1, proofBz, height1, err1 := tendermint.QueryTendermintProof(sourceClient.CoreSdk, int64(height.GetRevisionHeight()), key)

	fmt.Println("res", hex.EncodeToString(res.Commitment))
	fmt.Println("v1", hex.EncodeToString(v1))

	fmt.Println("height1", height1)
	fmt.Println("height", height.GetRevisionHeight())

	fmt.Println("proof ::!!!", proofBz)
	fmt.Println("height::!! ", height.GetRevisionHeight())
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	ress, err := destClient.RecvPacket(proofBz, packet1, int64(height.GetRevisionHeight()), res.ProofHeight.RevisionNumber, baseTx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ress)
}
