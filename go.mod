module github.com/mapofzones/cosmos-watcher

go 1.15

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/cosmos/cosmos-sdk => github.com/mapofzones/cosmos-sdk v0.42.9-unmarshal-fix

require (
	github.com/cosmos/cosmos-sdk v0.45.12
	github.com/gogo/protobuf v1.3.3
	github.com/medibloc/panacea-core/v2 v2.0.6
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.8.1
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.34.24
)
