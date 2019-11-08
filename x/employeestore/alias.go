package employeestore

import (
	"github.com/sahith-narahari/sdk-tutorial/x/employeestore/internal/keeper"

	"github.com/sahith-narahari/sdk-tutorial/x/employeestore/internal/types"
)

type (
	Keeper = keeper.Keeper
	//RegisterCodec = types.RegisterCodec
	MsgStoreName   = types.MsgEmployee
	MsgGetEmployee = types.MsgEmployee
)

const (
	ModuleName = types.ModuleName
)
