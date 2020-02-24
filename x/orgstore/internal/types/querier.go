package types

import "strings"

// Query Result Payload for a resolve query
type QueryResResolve struct {
	Value string `json:"value"`
}

// implement fmt.Stringer
func (r QueryResResolve) String() string {
	return r.Value
}

// Query Result Payload for a names query
type QueryResNames []string

// implement fmt.Stringer
func (n QueryResNames) String() string {
	return strings.Join(n[:], "\n")
}

// Query Result Payload for a resolve query
type QueryOrgResolve struct {
	Name string `json:"name"`
	Address string `json:"address"`
	Employees QueryUsersResolve `json:"employees,omitempty"`
}

type QueryOrgsResolve []QueryOrgResolve

// implement fmt.Stringer
func (r QueryOrgResolve) String() string {
	return r.Name
}

func (r QueryOrgsResolve) String() string {
return ""
}

type QueryUserResolve struct {
	Name string `json:"name"`
	Address string `json:"address"`
}

type QueryUsersResolve []QueryUserResolve

func (r QueryUserResolve) String() string {
	return ""
}

func (r QueryUsersResolve) String() string {
	return ""
}
