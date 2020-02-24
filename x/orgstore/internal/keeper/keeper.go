package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"nameservice/x/orgstore/internal/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	CoinKeeper bank.Keeper

	storeKey  sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// Organization store

func (k Keeper) SetOrg(ctx sdk.Context,companyName string,address string)  {
	store := ctx.KVStore(k.storeKey)
	org := types.Org{
		Name:companyName,
		Address:address,
	}
	store.Set([]byte(companyName),k.cdc.MustMarshalBinaryBare(org))
}

func (k Keeper) GetUserName(ctx sdk.Context,orgName string, userName string) types.Employee {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(orgName))
	var org types.Org
	k.cdc.MustUnmarshalBinaryBare(bz,&org)
	for _, v := range org.Employees {
		if v.Name == userName {
			return v
		}
	}

	return types.Employee{}
}

func (k Keeper) SetUser(ctx sdk.Context,name string,address string,org string)  {
	store := ctx.KVStore(k.storeKey)
	org1 := k.GetOrg(ctx,org)
	emp :=types.Employee{
		Name: name,
		Address:   address,
	}

	org1.Employees = append(org1.Employees, emp)
	store.Set([]byte(org),k.cdc.MustMarshalBinaryBare(org1))
}

func (k Keeper) GetOrg(ctx sdk.Context, name string) types.Org {
	store := ctx.KVStore(k.storeKey)
	if !k.IsNamePresent(ctx, name) {
		return types.Org{}
	}

	bz := store.Get([]byte(name))
	var org types.Org
	k.cdc.MustUnmarshalBinaryBare(bz,&org)
	return org
}


func (k Keeper) GetOrgName(ctx sdk.Context, name string) string {
	return k.GetOrg(ctx, name).Name
}

func (k Keeper) DeleteOrg(ctx sdk.Context, name string) {
		store := ctx.KVStore(k.storeKey)
		store.Delete([]byte(name))
}

// Check if the name is present in the store or not
func (k Keeper) IsNamePresent(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(name))
}

func (k Keeper) GetOrgIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}

// NewKeeper creates new instances of the orgstore Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		CoinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}