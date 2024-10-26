package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddAccountInfo{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdatetotalshares{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInitiateDeathValidation{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgValidateDeath{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMultisite{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateMultisite{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateEmailId{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
