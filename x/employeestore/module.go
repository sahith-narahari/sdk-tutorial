package employeestore

import (
	"github.com/comdex-blockchain/x/bank"
	"github.com/cosmos-cg-key-management/cosmos-sdk/types/module"
	//"github.com/cosmos/cosmos-sdk/codec"
)

var (
	_ module.AppModule = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

type AppModuleBasic struct{}

type AppModule struct {
	AppModuleBasic
	keeper     Keeper
	coinKeeper bank.Keeper
}

func (AppModuleBasic) Name() string {
	return ModuleName
}

//func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
//	RegisterCodec(cdc)
//}