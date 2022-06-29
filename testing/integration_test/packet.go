package integration

import (
	"fmt"
	"strconv"

	tibc "github.com/bianjieai/tibc-sdk-go"
	"github.com/bianjieai/tibc-sdk-go/modules/core/packet"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types"
)

func queryCommitment(client tibc.Client, destName string, sourceName string, seq uint64) *packet.QueryPacketCommitmentResponse {
	res, err := client.PacketCommitment(destName, sourceName, seq)
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
	return packet.Packet{
		Sequence:         uint64(num),
		SourceChain:      sourceChain,
		DestinationChain: destinationChain,
		Port:             port,
		RelayChain:       relayChain,
		Data:             []byte(data),
	}, []byte(ack), nil
}
func sendAck(sourceClient Client, destClient Client, keyname string, txhash string) (string, tibctypes.IError) {
	tx, err := destClient.QueryTx(txhash)
	if err != nil {
		return "", tibctypes.New("querytx", 0, "error query tx")
	}
	clients, err := sourceClient.Tendermint.GetClientState(destClient.ChainName)
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
	_, proofBz, _, err1 := destClient.Tendermint.QueryTendermintProof(int64(height.GetRevisionHeight()), key)

	if err1 != nil {
		fmt.Println(err1)
		return "", tibctypes.New("queryProof", 0, "error query proof")

	}
	ress, err := sourceClient.Tendermint.Acknowledgement(proofBz, ack, packet1, int64(height.GetRevisionHeight()), clients.GetLatestHeight().GetRevisionNumber(), baseTx)
	if err != nil {
		fmt.Println(err)
		return "", tibctypes.New("queryProof", 0, "error send acknowledgement")
	}
	fmt.Println(ress)
	return ress.Hash, nil
}

func packetRecive(sourceClient Client, destClient Client, keyname string, txHash string) (string, tibctypes.IError) {
	tx, err := sourceClient.QueryTx(txHash)
	if err != nil {
		fmt.Println(err)
		return "", tibctypes.New("querytx", 0, "error query tx")
	}
	clients, err := destClient.Tendermint.GetClientState(sourceClient.ChainName)

	height := clients.GetLatestHeight()
	packet1, err := getpacket(tx)
	//fmt.Println(packet1.String())
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
	_, proofBz, _, err1 := sourceClient.Tendermint.QueryTendermintProof(int64(height.GetRevisionHeight()), key)
	if err1 != nil {
		fmt.Println(err1)
		return "", tibctypes.New("queryProof", 0, "error query proof")
	}
	ress, err := destClient.Tendermint.RecvPacket(proofBz, packet1, int64(height.GetRevisionHeight()), clients.GetLatestHeight().GetRevisionNumber(), baseTx)
	if err != nil {
		return "", tibctypes.New("recvPacket", 0, "error recive packet")
	}
	fmt.Println(ress)
	return ress.Hash, nil
}
func cleanPacket(sourceClient, destClient Client, seq uint64, keyname string) (string, tibctypes.IError) {
	cleanpacket := packet.CleanPacket{
		Sequence:         seq,
		SourceChain:      sourceClient.ChainName,
		DestinationChain: destClient.ChainName,
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
	res, err := sourceClient.Tendermint.CleanPacket(cleanpacket, baseTx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(res)
	return res.Hash, nil
}
func recvCleanPacket(sourceClient, destClient Client, keyname string, txhash string) (string, tibctypes.IError) {
	tx, err := sourceClient.QueryTx(txhash)
	if err != nil {
		fmt.Println(err)
		return "", tibctypes.New("querytx", 0, "error query tx")
	}
	clients, err := destClient.Tendermint.GetClientState(sourceClient.ChainName)
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
	_, proofBz, _, err1 := sourceClient.Tendermint.QueryTendermintProof(int64(height.GetRevisionHeight()), key)
	if err1 != nil {
		fmt.Println(err1)
		return "", tibctypes.New("queryProof", 0, "error query proof")
	}
	ress, err := destClient.Tendermint.RecvCleanPacket(proofBz, cleanpack, int64(height.GetRevisionHeight()), clients.GetLatestHeight().GetRevisionNumber(), baseTx)
	if err != nil {
		fmt.Println(err)
		return "", tibctypes.New("recvcleanpacket", 0, "error Recv CleanPacket")
	}
	fmt.Println(ress)
	return ress.Hash, nil
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
	return packet.CleanPacket{
		Sequence:         uint64(num),
		SourceChain:      sourceChain,
		DestinationChain: destinationChain,
		RelayChain:       relayChain,
	}, nil

}
