package organizationStore

import (
	"github.com/sahith-narahari/sdk-tutorial/x/organizationStore/internal/keeper"
	"github.com/sahith-narahari/sdk-tutorial/x/organizationStore/internal/types"
)

type (
	Keeper                 = keeper.Keeper
	MsgStoreOrganization   = types.MsgOrgStore
	MsgStoreOrgUser        = types.OrgUsers
	MsgDeleteOrganizations = types.MsgDeleteOrganization
	MsgDeleteOrgUser       = types.MsgDeleteOrgUser
)

const (
	ModuleName = types.ModuleName
	StoreKey   = ModuleName
	RouterKey  = ModuleName
)

var (
	RegisterCodec = types.RegisterCodec
	ModuleCdc     = types.ModuleCdc
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier
)
