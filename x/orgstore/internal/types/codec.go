package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	//cdc.RegisterConcrete(MsgSetName{}, "orgstore/SetName", nil)
	cdc.RegisterConcrete(OrgSet{}, "orgstore/SetOrg", nil)
	cdc.RegisterConcrete(UserSet{}, "orgstore/SetUser", nil)
	cdc.RegisterConcrete(MsgDeleteOrg{}, "orgstore/DeleteName", nil)
	//cdc.RegisterConcrete(MsgBuyName{}, "orgstore/BuyName", nil)
	//cdc.RegisterConcrete(MsgDeleteName{}, "orgstore/DeleteName", nil)
}
