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
	res, err := client.PacketCommitment("testCreateClientB", "testCreateClientA", 1)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(res.String())
	return res
}

func packetReceipt(client tibc.Client) {
	res, err := client.PacketReceipt("testCreateClientB", "testCreateClientA", 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Received)
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

func packetRecive(sourceClient tibc.Client, destClient tibc.Client, keyname string) {
	tx, err := sourceClient.CoreSdk.QueryTx("D6C9C31731F54D0D98CF93538678B03F4E0A10F43B23C8B3EA7A5394CEC256A1")
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
