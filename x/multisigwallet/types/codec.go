package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateWallet{}, "multisigwallet/CreateWallet", nil)
	cdc.RegisterConcrete(&MsgCreateTransaction{}, "multisigwallet/CreateTransaction", nil)
	cdc.RegisterConcrete(&MsgSignTransaction{}, "multisigwallet/SignTransaction", nil)
	cdc.RegisterConcrete(&MsgExecuteTransaction{}, "multisigwallet/ExecuteTransaction", nil)
	cdc.RegisterConcrete(&MsgAddSigners{}, "multisigwallet/AddSigners", nil)
	cdc.RegisterConcrete(&MsgRemoveSigners{}, "multisigwallet/RemoveSigners", nil)
	cdc.RegisterConcrete(&MsgUpdateMinSigns{}, "multisigwallet/UpdateMinSigns", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateWallet{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTransaction{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSignTransaction{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgExecuteTransaction{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddSigners{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveSigners{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateMinSigns{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)