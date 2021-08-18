module github.com/bianjieai/tibc-sdk-go

go 1.15

require (
	github.com/confio/ics23/go v0.6.6
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/irisnet/core-sdk-go v0.0.0-20210817104504-bd2c112847e9
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/tendermint/tendermint v0.34.11
	google.golang.org/genproto v0.0.0-20210805201207-89edb61ffb67
	google.golang.org/grpc v1.39.1
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)
