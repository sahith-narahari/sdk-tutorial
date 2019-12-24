package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgOrgStore struct {
	Name     string         `json:"name"`
	Owner    sdk.AccAddress `json:"owner"`
	OrgUsers []OrgUsers     `json:"orgUsers"`
}

func NewMsgOrganizationStore(name string, owner sdk.AccAddress, uName string, uRole string, id string) MsgOrgStore {
	org := MsgOrgStore{
		Name:  name,
		Owner: owner,
	}
	return org
}

func (msg MsgOrgStore) Type() string {
	return "set_org_store"
}

func (msg MsgOrgStore) Route() string {
	return RouterKey
}

func (msg MsgOrgStore) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}

	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Organization name can't be empty")
	}

	return nil
}

func (msg MsgOrgStore) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgOrgStore) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

type OrgUsers struct {
	Address sdk.AccAddress `json:"address"`
	Name    string         `json:"name"`
	Role    string         `json:"role"`
	OrgName string         `json:"orgName"`
}

func NewMsgOrganizationUserStore(orgName string, owner sdk.AccAddress, uName string, uRole string) OrgUsers {

	OrgUsers := OrgUsers{
		OrgName: orgName,
		Address: owner,
		Name:    uName,
		Role:    uRole,
	}

	return OrgUsers
}

func (msg OrgUsers) Type() string {
	return "set_org_user"
}

func (msg OrgUsers) Route() string {
	return RouterKey
}

func (msg OrgUsers) ValidateBasic() sdk.Error {
	if msg.Address.Empty() {
		return sdk.ErrInvalidAddress(msg.Address.String())
	}

	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("User name can't be empty")
	}

	return nil
}

func (msg OrgUsers) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg OrgUsers) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}

type MsgDeleteOrganization struct {
	Name     string         `json:"name"`
	Owner    sdk.AccAddress `json:"owner"`
	OrgUsers []OrgUsers     `json:"orgUsers"`
}

func NewMsgDeleteOrganization(name string, owner sdk.AccAddress) MsgDeleteOrganization {
	org := MsgDeleteOrganization{
		Name:  name,
		Owner: owner,
	}
	return org
}

func (msg MsgDeleteOrganization) Type() string {
	return "delete_organization"
}

func (msg MsgDeleteOrganization) Route() string {
	return RouterKey
}

func (msg MsgDeleteOrganization) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}

	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Organization name can't be empty")
	}

	return nil
}

func (msg MsgDeleteOrganization) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgDeleteOrganization) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

type MsgDeleteOrgUser struct {
	Address sdk.AccAddress `json:"address"`
	Name    string         `json:"name"`
	Role    string         `json:"role"`
	OrgName string         `json:"orgName"`
}

func NewMsgDeleteOrgUser(orgName string, owner sdk.AccAddress, userName string) MsgDeleteOrgUser {
	org := MsgDeleteOrgUser{
		Name:    userName,
		Address: owner,
		OrgName:orgName,
	}
	return org
}

func (msg MsgDeleteOrgUser) Type() string {
	return "delete_org_user"
}

func (msg MsgDeleteOrgUser) Route() string {
	return RouterKey
}

func (msg MsgDeleteOrgUser) ValidateBasic() sdk.Error {
	if msg.Address.Empty() {
		return sdk.ErrInvalidAddress(msg.Address.String())
	}

	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("user name can't be empty")
	}

	return nil
}

func (msg MsgDeleteOrgUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgDeleteOrgUser) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
