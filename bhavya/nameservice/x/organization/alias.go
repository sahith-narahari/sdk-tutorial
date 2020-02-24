package orgs

import (
	"github.com/cosmos/bhavya/nameservice/x/organization/internal/keeper"
	"github.com/cosmos/bhavya/nameservice/x/organization/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper       = keeper.NewKeeper
	NewQuerier      = keeper.NewQuerier
	NewMsgSetOrg    = types.NewMsgSetOrg
	NewMsgUpdateOrg = types.NewMsgUpdateOrg
	ModuleCdc       = types.ModuleCdc
	RegisterCodec   = types.RegisterCodec
)

type (
	Keeper       = keeper.Keeper
	MsgSetOrg    = types.MsgSetOrg
	MsgUpdateOrg = types.MsgUpdateOrg
	Orgs         = types.Org
	MsgDeleteOrg = types.MsgDeleteOrg
)
