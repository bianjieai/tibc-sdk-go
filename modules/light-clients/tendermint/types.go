package tendermint

type EthermintConfig struct {
	Prefix          []byte `json:"prefix"`
	ContractAddress string `json:"contract_address"`
}
