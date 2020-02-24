package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Transactions messages must fulfill the Msg
type Msg interface {
// Return the message type.
// Must be alphanumeric or empty.
Type() string

// Returns a human-readable string for the message, intended for utilization
// within tags
Route() string

// ValidateBasic does a simple validation check that
// doesn't require access to any other information.
ValidateBasic() sdk.Error

// Get the canonical byte representation of the Msg.
GetSignBytes() []byte

// Signers returns the addrs of signers that must sign.
// CONTRACT: All signatures must be present to be valid.
// CONTRACT: Returns addrs in some deterministic order.
GetSigners() []sdk.AccAddress
}

const RouterKey = ModuleName // this was defined in your key.go file

// MsgSetName defines a SetName message

type OrgSet struct {
	Address string `json:"address"`
	CompanyName string `json:"name"`
	Owner sdk.AccAddress `json:"owner"`
}



func NewMsgOrgSet(name string,address string,owner sdk.AccAddress) OrgSet {
	return OrgSet{
		Address:address,
		CompanyName: name,
		Owner:owner,
	}
}

func (msg OrgSet) Route() string { return RouterKey }

// Type should return the action
func (msg OrgSet) Type() string { return "set_org" }

func (msg OrgSet) ValidateBasic() sdk.Error {
	if len(msg.Address) == 0 || len(msg.CompanyName) == 0 {
		return sdk.ErrUnknownRequest("CEO and/or Company name cannot be empty")
	}
	return nil
}

func (msg OrgSet) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg OrgSet) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ msg.Owner }
}

type UserSet struct {
	Address string `json:"address"`
	Name string `json:"name"`
	Owner sdk.AccAddress `json:"owner"`
	Organization string `json:"organization"`
}


func NewMsgUserSet(org string,name string,address string,owner sdk.AccAddress) UserSet {
	return UserSet{
		Organization: org,
		Address:address,
		Name: name,
		Owner:owner,
	}
}

func (msg UserSet) Route() string { return RouterKey }

// Type should return the action
func (msg UserSet) Type() string { return "set_org" }

func (msg UserSet) ValidateBasic() sdk.Error {
	if len(msg.Address) == 0 || len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Employee name and/or address cannot be empty")
	}
	return nil
}

func (msg UserSet) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg UserSet) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{ msg.Owner }
}


type MsgDeleteOrg struct {
	Name  string         `json:"name"`
	Owner sdk.AccAddress `json:"owner"`
}

// NewMsgDeleteName is a constructor function for MsgDeleteName
func NewMsgDeleteName(name string, owner sdk.AccAddress) MsgDeleteOrg {
	return MsgDeleteOrg{
		Name:  name,
		Owner: owner,
	}
}

// Route should return the name of the module
func (msg MsgDeleteOrg) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeleteOrg) Type() string { return "delete_name" }

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteOrg) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteOrg) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteOrg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
