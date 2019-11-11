package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RouterKey is the module name router key
const RouterKey = ModuleName // this was defined in your key.go file

// MsgSetName defines a SetName message
type StoreEmp struct {
	Id 	  string 		`json:"id"`
	Name  string         `json:"name"`
	Owner sdk.AccAddress `json:"owner"`
}

// NewMsgSetName is a constructor function for MsgSetName
func NewEmpStore(name string, id string, owner sdk.AccAddress) StoreEmp {
	return StoreEmp{
		Name:  name,
		Id: id,
		Owner: owner,
	}
}

// Route should return the name of the module
func (msg StoreEmp) Route() string { return RouterKey }

// Type should return the action
func (msg StoreEmp) Type() string { return "set_name" }

// ValidateBasic runs stateless checks on the message
func (msg StoreEmp) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Id) == 0 {
		return sdk.ErrUnknownRequest("Name and/or id cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg StoreEmp) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg StoreEmp) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}