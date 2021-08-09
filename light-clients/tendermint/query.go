package tendermint

import (
	"github.com/bianjieai/tibc-sdk-go/client"
	grpc1 "github.com/gogo/protobuf/grpc"
)

type Client struct {
	client.QueryClient
}

func NewClient(cc grpc1.ClientConn) (*Client, error) {
	return &Client{
		QueryClient: client.NewQueryClient(cc),
	}, nil
}

func (q *Client) QueryClientState() (ClientState, error) {

	return ClientState{}, nil
}

func (q *Client) QueryConsensusState() (ConsensusState, error) {

	return ConsensusState{}, nil
}
