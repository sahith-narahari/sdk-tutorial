package types

import "github.com/cosmos/cosmos-sdk/codec"

var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgUpdateOrg{}, "orgs/UpdateOrg", nil)
	cdc.RegisterConcrete(MsgSetOrg{}, "orgs/SetOrg", nil)
	cdc.RegisterConcrete(MsgDeleteOrg{}, "orgs/DeleteOrg", nil)
}
