package keeper

import (
	"github.com/cosmos/bhavya/nameservice/x/organization/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	CoinKeeper bank.Keeper

	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// Sets the entire Org metadata struct for a name
func (k Keeper) SetOrg(ctx sdk.Context, name string, org types.Org) {

	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(org))
}

// updates the org employee
func (k Keeper) UpdateOrg(ctx sdk.Context, name string, id string, emp string) {
	org := k.GetOrg(ctx, name)
	org.Employee = append(org.Employee, emp)
	org.ID = id
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(org))
}

// Deletes the entire Org metadata struct for a name
func (k Keeper) DeleteOrg(ctx sdk.Context, name string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(name))
}

// Gets the entire Org metadata struct for a name
func (k Keeper) GetOrg(ctx sdk.Context, name string) types.Org {
	store := ctx.KVStore(k.storeKey)
	if isOrgPresent := store.Has([]byte(name)); !isOrgPresent {
		return types.Org{}
	}
	metadata := store.Get([]byte(name))
	var org types.Org
	k.cdc.MustUnmarshalBinaryBare(metadata, &org)
	return org
}

// NewKeeper creates new instances of the nameservice Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		CoinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

// Get an iterator over all names in which the keys are the names and the values are the whois
func (k Keeper) GetOrgsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)

	return sdk.KVStorePrefixIterator(store, []byte{})
}

// Check if the name is present in the store or not
func (k Keeper) IsOrgPresent(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(name))
}

// GetOwner - get the current owner of a name
func (k Keeper) GetOwner(ctx sdk.Context, name string) sdk.AccAddress {
	return k.GetOrg(ctx, name).Owner
}
