package studentstore

import (
	"github.com/sahith-narahari/sdk-tutorials/x/studentstore/internal/keeper"
	"github.com/sahith-narahari/sdk-tutorials/x/studentstore/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper = keeper.NewKeeper
	ModuleCdc    = types.ModuleCdc
	RegiserCodec = types.RegisterCodec

)

type (
	Keeper          = keeper.Keeper
	MsgStoreStudent = types.MsgStudentStore
)
