package orgstore

import (
"nameservice/x/orgstore/internal/keeper"
"nameservice/x/orgstore/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper        = keeper.NewKeeper
	NewQuerier       = keeper.NewQuerier
	ModuleCdc        = types.ModuleCdc
	RegisterCodec    = types.RegisterCodec
)

type (
	Keeper          = keeper.Keeper
	OrgSet      = types.OrgSet
	MsgDelete = types.MsgDeleteOrg
	UserSet      = types.UserSet
	QueryResResolve = types.QueryResResolve
	QueryResNames   = types.QueryResNames
	Org = types.Org
	Employee = types.Employee
)