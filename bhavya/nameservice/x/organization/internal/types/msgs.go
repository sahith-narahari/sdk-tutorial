package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName // this was defined in your key.go file

// MsgSetOrg defines a SetOrg message
type MsgSetOrg struct {
	Name  string         `json:"name"`
	CEO   string         `json:"ceo"`
	Owner sdk.AccAddress `json:"owner"`
}

// NewMsgSetOrg is a constructor function for MsgSetOrg
func NewMsgSetOrg(name string, ceo string, owner sdk.AccAddress) MsgSetOrg {
	return MsgSetOrg{
		Name:  name,
		CEO:   ceo,
		Owner: owner,
	}
}

// Route should return the name of the module
func (msg MsgSetOrg) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetOrg) Type() string { return "set_org" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetOrg) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.CEO) == 0 {
		return sdk.ErrUnknownRequest("Org Name and/or CEO cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetOrg) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetOrg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgSetOrg defines a SetOrg message
type MsgUpdateOrg struct {
	Name     string         `json:"name"`
	Employee string         `json:"employee"`
	ID       string         `json:"id"`
	Owner    sdk.AccAddress `json:"owner"`
}

// NewMsgSetOrg is a constructor function for MsgSetOrg
func NewMsgUpdateOrg(name string, emp string, id string, address sdk.AccAddress) MsgUpdateOrg {
	return MsgUpdateOrg{
		Name:     name,
		Employee: emp,
		ID:       id,
		Owner:    address,
	}
}

// Route should return the name of the module
func (msg MsgUpdateOrg) Route() string {
	return RouterKey
}

// Type should return the action
func (msg MsgUpdateOrg) Type() string {
	return "update_org"
}

// ValidateBasic runs stateless checks on the msg
func (msg MsgUpdateOrg) ValidateBasic() sdk.Error {
	if len(msg.Employee) == 0 || len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Org Name and/or Employee name cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgUpdateOrg) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgUpdateOrg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgDeleteOrg defines a SetOrg message
type MsgDeleteOrg struct {
	Name     string         `json:"name"`
	Owner    sdk.AccAddress `json:"owner"`
}

// NewMsgSetOrg is a constructor function for MsgSetOrg
func NewMsgDeleteOrg(name string, address sdk.AccAddress) MsgDeleteOrg {
	return MsgDeleteOrg{
		Name:     name,
		Owner:    address,
	}
}

// Route should return the name of the module
func (msg MsgDeleteOrg) Route() string {
	return RouterKey
}

// Type should return the action
func (msg MsgDeleteOrg) Type() string {
	return "update_org"
}

// ValidateBasic runs stateless checks on the msg
func (msg MsgDeleteOrg) ValidateBasic() sdk.Error {
	if len(msg.Owner) == 0 || len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Org Name and/or Owner name cannot be empty")
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
