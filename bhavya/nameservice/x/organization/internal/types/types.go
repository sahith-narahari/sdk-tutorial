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

//func NewOrg() Org {
//	return Org{}
//}
//
//// implement fmt.Stringer
//func (o Org) String() string {
//	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
//Name: %s
//CEO: %s`, o.Owner, o.Name, o.CEO))
//}
