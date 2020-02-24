package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type QueryResOrg struct {
	Name     string         `json:"name"`
	CEO      string         `json:"ceo"`
	Owner    sdk.AccAddress `json:"owner"`
	Employee []string       `json:"employee"`
}

func (r QueryResOrg) String() string {
	return fmt.Sprintln("Name: ", r.Name, "\nCEO: ", r.CEO, "\nOwner", r.Owner)
}

type QueryResOrgs []QueryResOrg

func (r QueryResOrgs) String() string {
	return ""
}
