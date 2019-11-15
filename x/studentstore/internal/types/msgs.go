package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName


type MsgStudentStore struct {
	Id 	string 		`json:"id"`
	Name string 	`json:"name"`
	Owner sdk.AccAddress  `json:"owner"`
}

func NewMsgStudentStore(id string, name string, owner sdk.AccAddress) MsgStudentStore  {
	return MsgStudentStore{
		Id:id,
		Name: name,
		Owner: owner,
	}
}

func (msg MsgStudentStore) Type() string {
	return "set_student_store"
}

func (msg MsgStudentStore) Route() string  {
	return RouterKey
}

func (msg MsgStudentStore) ValidateBasic() sdk.Error  {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}

	if len(msg.Name) == 0 || len(msg.Id) == 0 {
		return sdk.ErrUnknownRequest("Name/ Id can't be empty")
	}

	return nil
}

func (msg MsgStudentStore) GetSignBytes()[]byte  {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgStudentStore) GetSigners() []sdk.AccAddress  {
	return []sdk.AccAddress{msg.Owner}
}