package packet

import "fmt"

// PacketCommitmentKey returns the store key of under which a packet commitment
// is stored
func PacketCommitmentKey(sourceChain, destinationChain string, sequence uint64) []byte {
	return []byte(PacketCommitmentPath(sourceChain, destinationChain, sequence))
}

// PacketCommitmentPath defines the commitments to packet data fields store path
func PacketCommitmentPath(sourceChain, destinationChain string, sequence uint64) string {
	return fmt.Sprintf("%s/%d", PacketCommitmentPrefixPath(sourceChain, destinationChain), sequence)
}

// PacketCommitmentPrefixPath defines the prefix for commitments to packet data fields store path.
func PacketCommitmentPrefixPath(sourceChain, destinationChain string) string {
	return fmt.Sprintf("%s/%s/%s", "commitments", packetPath(sourceChain, destinationChain), "sequences")
}

func packetPath(sourceChain, destinationChain string) string {
	return fmt.Sprintf("%s/%s", sourceChain, destinationChain)
}