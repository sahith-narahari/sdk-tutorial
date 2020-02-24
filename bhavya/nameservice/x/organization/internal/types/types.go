package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Org struct {
	Name     string         `json:"name"`
	CEO      string         `json:"ceo"`
	Owner    sdk.AccAddress `json:"owner"`
	Employee []string       `json:"employee"`
	ID       string         `json:"id"`
}
