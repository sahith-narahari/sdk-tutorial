package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

)

const RouterKey = ModuleName

type MsgEmployee struct {
	Name       string         `json:"name"`
	EmployeeId string         `json:"employeeId"`
}

type MsgGetEmployee struct {
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgStoreEmployee(name string, id string) MsgEmployee {
	return MsgEmployee{
		Name:       name,
		EmployeeId: id,
	}
}

// Route should return the name of the module
func (msg MsgEmployee) Route() string { return RouterKey }

// Type should return the action
func (msg MsgEmployee) Type() string { return "store_employee" }

func (msg MsgEmployee) ValidateBasic() sdk.Error {
	if len(msg.Name) == 0 || len(msg.EmployeeId) == 0 {
		return sdk.ErrUnknownRequest("Name/Id cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgEmployee) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}