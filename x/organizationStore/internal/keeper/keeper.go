package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sahith-narahari/sdk-tutorial/x/organizationStore/internal/types"
)

type Keeper struct {
	Storekey sdk.StoreKey
	cdc      *codec.Codec
}

func NewKeeper(storekey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		storekey,
		cdc,
	}
}

func (k Keeper) SetOrganization(ctx sdk.Context, Name string, owner sdk.AccAddress) {
	var org types.MsgOrgStore

	if owner.Empty() {
		return
	}

	org.Name = Name
	org.Owner = owner

	store := ctx.KVStore(k.Storekey)
	store.Set([]byte(Name), k.cdc.MustMarshalBinaryBare(org))
}

func (k Keeper) SetOrganizationUser(ctx sdk.Context, orgName string, owner sdk.AccAddress, role string, uName string) {
	var org types.MsgOrgStore

	if owner.Empty() {
		return
	}
	orgUser := types.OrgUsers{
		Address: owner,
		Role:    role,
		Name:    uName,
		OrgName: orgName,
	}
	var user []types.OrgUsers

	org.OrgUsers = append(user, orgUser)

	store := ctx.KVStore(k.Storekey)
	store.Set([]byte(orgName), k.cdc.MustMarshalBinaryBare(org))
}

func (k Keeper) GetOrganization(ctx sdk.Context, name string) types.MsgOrgStore {
	var orgStore types.MsgOrgStore
	store := ctx.KVStore(k.Storekey)

	bz := store.Get([]byte(name))

	if bz == nil {
		return orgStore
	}

	k.cdc.MustUnmarshalBinaryBare(bz, &orgStore)
	return orgStore
}

// Check if the name is present in the store or not
func (k Keeper) IsNamePresent(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.Storekey)
	return store.Has([]byte(name))
}

func (k Keeper) GetWhichOrg(ctx sdk.Context, name string) types.MsgDeleteOrganization {
	store := ctx.KVStore(k.Storekey)
	if !k.IsNamePresent(ctx, name) {
		return types.MsgDeleteOrganization{}
	}
	bz := store.Get([]byte(name))
	var org types.MsgDeleteOrganization
	k.cdc.MustUnmarshalBinaryBare(bz, &org)
	return org
}

func (k Keeper) GetOwner(ctx sdk.Context, name string) sdk.AccAddress {
	return k.GetWhichOrg(ctx, name).Owner
}

func (k Keeper) DeleteOrganization(ctx sdk.Context, name string) {
	store := ctx.KVStore(k.Storekey)
	store.Delete([]byte(name))
}

func (k Keeper) DeleteOrgUser(ctx sdk.Context, orgName string, uName string) {
	store := ctx.KVStore(k.Storekey)

	d := store.Get([]byte(orgName))
	var org types.MsgDeleteOrgUser
	k.cdc.MustUnmarshalBinaryBare(d, &org)

	if !k.IsNamePresent(ctx, uName) {
		return
	} else {
		store.Delete([]byte(uName))
	}

	return
}

func (k Keeper) DeleteOrganizationUser(ctx sdk.Context, orgName string, uName string) {
	store := ctx.KVStore(k.Storekey)
	orgDetails := k.GetOrganization(ctx, orgName)

	var details []types.OrgUsers
	for _, org := range orgDetails.OrgUsers {
		if org.Name != uName {
			details = append(details, org)
		}
	}

	store.Set([]byte(orgName), k.cdc.MustMarshalBinaryBare(details))
}


