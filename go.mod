module github.com/bianjieai/tibc-sdk-go

go 1.15

require (
	github.com/DataDog/zstd v1.4.5 // indirect
	github.com/confio/ics23/go v0.6.6
	github.com/dgraph-io/ristretto v0.0.3 // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/ethereum/go-ethereum v1.10.7
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/irisnet/core-sdk-go v0.0.0-20210917071031-be43063a1899
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.10.0 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/tendermint/tendermint v0.34.11
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2
	google.golang.org/genproto v0.0.0-20210805201207-89edb61ffb67
	google.golang.org/grpc v1.39.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)
