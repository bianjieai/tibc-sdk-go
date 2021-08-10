package tendermint

import (
	"github.com/bianjieai/tibc-sdk-go/client"
)

func (q *tendermintLightClient) UpdateClient(msgUpdateClient client.UpdateClientRequest) (*client.MsgUpdateClientResponse, error) {
	// todo? change Response?
	return q.cli.UpdateClient(msgUpdateClient)
}
