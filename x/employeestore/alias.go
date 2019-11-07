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
	NewKeeper = keeper.NewKeeper

)

type (
	Keeper          = keeper.Keeper

)

