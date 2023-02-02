package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateSubspace{}, "huddle/MsgCreateSubspace", nil)
	cdc.RegisterConcrete(&MsgEditSubspace{}, "huddle/MsgEditSubspace", nil)
	cdc.RegisterConcrete(&MsgCreateUserGroup{}, "huddle/MsgCreateUserGroup", nil)

	cdc.RegisterConcrete(&MsgCreateSection{}, "huddle/MsgCreateSection", nil)
	cdc.RegisterConcrete(&MsgEditSection{}, "huddle/MsgEditSection", nil)
	cdc.RegisterConcrete(&MsgMoveSection{}, "huddle/MsgMoveSection", nil)
	cdc.RegisterConcrete(&MsgDeleteSection{}, "huddle/MsgDeleteSection", nil)

	cdc.RegisterConcrete(&MsgEditUserGroup{}, "huddle/MsgEditUserGroup", nil)
	cdc.RegisterConcrete(&MsgMoveUserGroup{}, "huddle/MsgMoveUserGroup", nil)
	cdc.RegisterConcrete(&MsgSetUserGroupPermissions{}, "huddle/MsgSetUserGroupPermissions", nil)
	cdc.RegisterConcrete(&MsgDeleteUserGroup{}, "huddle/MsgDeleteUserGroup", nil)

	cdc.RegisterConcrete(&MsgAddUserToUserGroup{}, "huddle/MsgAddUserToUserGroup", nil)
	cdc.RegisterConcrete(&MsgRemoveUserFromUserGroup{}, "huddle/MsgRemoveUserFromUserGroup", nil)

	cdc.RegisterConcrete(&MsgSetUserPermissions{}, "huddle/MsgSetUserPermissions", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSubspace{},
		&MsgEditSubspace{},
		&MsgCreateUserGroup{},
		&MsgCreateSection{},
		&MsgEditSection{},
		&MsgMoveSection{},
		&MsgDeleteSection{},
		&MsgEditUserGroup{},
		&MsgMoveUserGroup{},
		&MsgSetUserGroupPermissions{},
		&MsgDeleteUserGroup{},
		&MsgAddUserToUserGroup{},
		&MsgRemoveUserFromUserGroup{},
		&MsgSetUserPermissions{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// AminoCodec references the global x/subspaces module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/subspaces and
	// defined at the application level.
	AminoCodec = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
}
