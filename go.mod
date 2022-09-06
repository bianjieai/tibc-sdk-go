module github.com/bianjieai/tibc-sdk-go

go 1.15

require (
	github.com/confio/ics23/go v0.6.6
	github.com/ethereum/go-ethereum v1.10.16
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/irisnet/core-sdk-go v0.0.0-20220906070548-0c9d0a868f37
	github.com/pkg/errors v0.9.1
	github.com/tendermint/tendermint v0.34.21
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e
	google.golang.org/genproto v0.0.0-20220725144611-272f38e5d71b
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.0
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.21-irita-220906
)
