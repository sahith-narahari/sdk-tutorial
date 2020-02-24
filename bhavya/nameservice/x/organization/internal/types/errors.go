package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultCodespace is the Module Name
const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeOrgDoesNotExist sdk.CodeType = 101
)

// ErrNameDoesNotExist is the error for name not existing
func ErrOrgDoesNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeOrgDoesNotExist, "Org does not exist")
}