module github.com/bianjieai/tibc-sdk-go

go 1.15

require (
	github.com/armon/go-metrics v0.3.9 // indirect
	github.com/confio/ics23/go v0.6.6
	github.com/cosmos/cosmos-sdk v0.42.9
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/irisnet/core-sdk-go v0.0.0-20210729072452-06544f6270f3
	github.com/spf13/cast v1.4.0 // indirect
	github.com/spf13/cobra v1.2.1 // indirect
	github.com/tendermint/tendermint v0.34.11
	google.golang.org/genproto v0.0.0-20210805201207-89edb61ffb67
	google.golang.org/grpc v1.39.1
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)
