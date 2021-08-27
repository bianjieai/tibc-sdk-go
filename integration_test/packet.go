package integration

import (
	"fmt"
	"strconv"

	tibc "github.com/bianjieai/tibc-sdk-go"
	"github.com/bianjieai/tibc-sdk-go/packet"
	"github.com/bianjieai/tibc-sdk-go/tendermint"
	"github.com/irisnet/core-sdk-go/types"
)

func queryCommitment(client tibc.Client, destName string, sourceName string) *packet.QueryPacketCommitmentResponse {
	res, err := client.PacketCommitment(destName, sourceName, 1)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(res.String())
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
func sendAck(sourceClient tibc.Client, destClient tibc.Client, keyname string, sourceName string) {
	tx, err := sourceClient.CoreSdk.QueryTx("CBF30E33349768EC9D32AD8C79D8C0B13E671F242E7E7A5783AB62DEF51CB478")
	if err != nil {
		fmt.Println(err)
		return
	}

	clients, err := destClient.GetClientState(sourceName)
	height := clients.GetLatestHeight()
	packet1, ack, err := getpacketAndAck(tx)
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
	_, proofBz, _, err1 := tendermint.QueryTendermintProof(sourceClient.CoreSdk, int64(height.GetRevisionHeight()), key)

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

func packetRecive(sourceClient tibc.Client, destClient tibc.Client, keyname string, sourceName string) {
	tx, err := sourceClient.CoreSdk.QueryTx("2E1687540F3E3A7BCDE793154B8FEB0F5CA7111074D21EE6FB8609BFEDB8E0B5")
	if err != nil {
		fmt.Println(err)
		return
	}
	res := queryCommitment(sourceClient, "testCreateClientB", sourceName)
	clients, err := destClient.GetClientState("testCreateClientA")
	height := clients.GetLatestHeight()
	packet1, err := getpacket(tx)
	fmt.Println(packet1.String())
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
	_, proofBz, _, err1 := tendermint.QueryTendermintProof(sourceClient.CoreSdk, int64(height.GetRevisionHeight()), key)
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
func cleanPacket(sourceClient tibc.Client, keyname string) {
	cleanpacket := packet.CleanPacket{
		Sequence:         1,
		SourceChain:      "testCreateClientA",
		DestinationChain: "testCreateClientC",
		RelayChain:       "",
	}
	baseTx := types.BaseTx{
		From:               keyname,
		Gas:                0,
		Memo:               "TEST",
		Mode:               types.Commit,
		Password:           "12345678",
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}
	res, err := sourceClient.CleanPacket(cleanpacket, baseTx)
	if err != nil {
		fmt.Println(err.Error(), err.Codespace())
		return
	}
	fmt.Println(res)

}
func recvCleanPacket(sourceClient tibc.Client, destClient tibc.Client, keyname string) {
	tx, err := sourceClient.CoreSdk.QueryTx("ECDD26B95971537A089E9DB1E02EB30B5E433281063504614D12A093686E6B4A")
	if err != nil {
		fmt.Println(err)
		return
	}
	clients, err := destClient.GetClientState("testCreateClientA")
	height := clients.GetLatestHeight()
	cleanpack, err := getcleanpack(tx)
	if err != nil {
		fmt.Println("getcleanpack error")
	}
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
	key := packet.CleanPacketCommitmentKey(cleanpack.GetSourceChain(), cleanpack.GetDestChain())
	_, proofBz, _, err1 := tendermint.QueryTendermintProof(sourceClient.CoreSdk, int64(height.GetRevisionHeight()), key)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	ress, err := destClient.RecvCleanPacket(proofBz, cleanpack, int64(height.GetRevisionHeight()), 0, baseTx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ress)
}

func getcleanpack(tx types.ResultQueryTx) (packet.CleanPacket, error) {
	sequence, err := tx.Result.Events.GetValue("send_clean_packet", "packet_sequence")
	if err != nil {
		fmt.Println(err)
		return packet.CleanPacket{}, nil
	}
	sourceChain, err := tx.Result.Events.GetValue("send_clean_packet", "packet_src_chain")
	if err != nil {
		fmt.Println(err)
		return packet.CleanPacket{}, nil
	}
	destinationChain, err := tx.Result.Events.GetValue("send_clean_packet", "packet_dst_port")
	if err != nil {
		fmt.Println(err)
		return packet.CleanPacket{}, nil
	}
	relayChain, err := tx.Result.Events.GetValue("send_clean_packet", "packet_relay_channel")
	if err != nil {
		fmt.Println(err)
		return packet.CleanPacket{}, nil
	}
	num, err := strconv.Atoi(sequence)
	if err != nil {
		fmt.Println(err)
		return packet.CleanPacket{}, nil
	}
	//fmt.Println(num)
	return packet.CleanPacket{
		Sequence:         uint64(num),
		SourceChain:      sourceChain,
		DestinationChain: destinationChain,
		RelayChain:       relayChain,
	}, nil

}
