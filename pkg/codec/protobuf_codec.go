package watcher

import (
	"github.com/gogo/protobuf/proto"

	kiapp "github.com/KiFoundation/ki-tools/app"
	cosmoscodectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmoscryptoed "github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cosmoscryptomultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	cosmoscryptosecp "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cosmoscryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	cosmosibcexported "github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
	cosmosibcclients "github.com/cosmos/cosmos-sdk/x/ibc/light-clients/07-tendermint/types"
)

func RegisterInterfacesAndImpls(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	impls := getMessageImplementations()
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), impls...)
	kiRegisterInterfaces(interfaceRegistry)
	registerTypes(interfaceRegistry)
}

func kiRegisterInterfaces(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	kiapp.ModuleBasics.RegisterInterfaces(interfaceRegistry)
}

func registerTypes(interfaceRegistry cosmoscodectypes.InterfaceRegistry) { // todo: need to nest. Maybe we can remove it. Old code
	interfaceRegistry.RegisterInterface("cosmos.crypto.PubKey", (*cosmoscryptotypes.PubKey)(nil))
	interfaceRegistry.RegisterImplementations((*cosmoscryptotypes.PubKey)(nil), &cosmoscryptoed.PubKey{})
	interfaceRegistry.RegisterImplementations((*cosmoscryptotypes.PubKey)(nil), &cosmoscryptosecp.PubKey{})
	interfaceRegistry.RegisterImplementations((*cosmoscryptotypes.PubKey)(nil), &cosmoscryptomultisig.LegacyAminoPubKey{})

	interfaceRegistry.RegisterImplementations((*cosmosibcexported.ClientState)(nil), &cosmosibcclients.ClientState{})
	interfaceRegistry.RegisterImplementations((*cosmosibcexported.ConsensusState)(nil), &cosmosibcclients.ConsensusState{})
	interfaceRegistry.RegisterImplementations((*cosmosibcexported.Header)(nil), &cosmosibcclients.Header{})
	interfaceRegistry.RegisterImplementations((*cosmosibcexported.Misbehaviour)(nil), &cosmosibcclients.Misbehaviour{})
}

func getMessageImplementations() []proto.Message {
	var impls []proto.Message
	cosmosMessages := getCosmosMessages()
	impls = append(impls, cosmosMessages...)
	return impls
}

func getCosmosMessages() []proto.Message {
	cosmosMessages := []proto.Message{
		&cosmostypes.ServiceMsg{}, // do i need it? cosmostypes.RegisterInterfaces don't exist ServiceMsg
	}
	return cosmosMessages
}
