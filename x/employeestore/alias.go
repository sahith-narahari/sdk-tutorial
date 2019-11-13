package employeestore

import (
	"github.com/cosmos/sdk-tutorials/nameservice/x/employeestore/internal/keeper"
	"github.com/cosmos/sdk-tutorials/nameservice/x/employeestore/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
)

type (
	Keeper       = keeper.Keeper
	MsgSetInfo   = types.MsgSetInfo
	EmployeeInfo = types.EmployeeInfo
)
