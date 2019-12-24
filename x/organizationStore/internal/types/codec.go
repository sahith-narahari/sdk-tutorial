package types


import "github.com/cosmos/cosmos-sdk/codec"

var ModuleCdc = codec.New()

func init()  {
	RegisterCodec(ModuleCdc)
}

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgOrgStore{}, "organizationStore/SetOrganization", nil)
	cdc.RegisterConcrete(MsgOrgStore{}, "organizationStore/SetOrganizationUser", nil)
	cdc.RegisterConcrete(MsgOrgStore{}, "organizationStore/DeleteOrganization", nil)
	cdc.RegisterConcrete(MsgOrgStore{}, "organizationStore/DeleteOrganizationUser", nil)
}