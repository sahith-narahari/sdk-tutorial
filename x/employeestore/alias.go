package employeestore

import (
	"github.com/sahith-narahari/sdk-tutorial/x/employeestore/internal/keeper"
	"github.com/sahith-narahari/sdk-tutorial/x/employeestore/internal/types"
)

type (
	Keeper         = keeper.Keeper
	MsgStoreName   = types.MsgStoreEmployee
	MsgGetEmployee = types.MsgStoreEmployee
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	RegisterCodec = types.RegisterCodec
	ModuleCdc     = types.ModuleCdc
	NewQuerier    = keeper.NewQuerier
)
