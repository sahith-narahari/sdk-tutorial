package types

import "github.com/cosmos/cosmos-sdk/codec"

var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgEmployee{}, "employeestore/Storename", nil)
	cdc.RegisterConcrete(MsgGetEmployee{},"employeestore/Getemployee", nil)
}