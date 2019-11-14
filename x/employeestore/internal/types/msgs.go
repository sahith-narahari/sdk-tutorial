package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName

var _ sdk.Msg = MsgStoreEmployee{}

type MsgStoreEmployee struct {
	Name       string         `json:"name"`
	EmployeeId string         `json:"employeeId"`
	Signers    sdk.AccAddress `json:"signers"`
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgStoreEmployee(name string, id string) MsgStoreEmployee {
	return MsgStoreEmployee{
		Name:       name,
		EmployeeId: id,
	}
}

// Route should return the name of the module
func (msg MsgStoreEmployee) Route() string { return RouterKey }

// Type should return the action
func (msg MsgStoreEmployee) Type() string { return "delete_name" }

// ValidateBasic runs stateless checks on the message
func (msg MsgStoreEmployee) ValidateBasic() sdk.Error {
	if msg.Signers.Empty() {
		return sdk.ErrInvalidAddress(msg.Signers.String())
	}
	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgStoreEmployee) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON)
}

// GetSigners defines whose signature is required
func (msg MsgStoreEmployee) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signers}
}
