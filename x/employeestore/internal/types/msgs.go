package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName

type MsgSetInfo struct {
	EmployeeInfo EmployeeInfo `json:"employee_name"`
}

func NewMsgSetInfo(EmployeeName string) MsgSetInfo {
	return MsgSetInfo{
		EmployeeInfo: EmployeeInfo{
			EmployeeName: EmployeeName,
		},
	}
}

func (msg MsgSetInfo) Route() string {
	return RouterKey
}

func (msg MsgSetInfo) Type() string {
	return "set_info"
}

func (msg MsgSetInfo) ValidateBasic() sdk.Error {
	if len(msg.EmployeeInfo.EmployeeName) == 0 {
		return sdk.ErrUnknownRequest("Name Cannot be empty")
	}
	return nil
}

func (msg MsgSetInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgSetInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}
