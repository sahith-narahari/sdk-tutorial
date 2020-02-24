package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"nameservice/x/orgstore/internal/types"
)

// query endpoints supported by the orgstore Querier
const (
	QueryOrg   = "org"
	QueryOrgs   = "orgs"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryOrg:
			return queryOrg(ctx,path,req,keeper)
		case QueryOrgs:
			return queryOrgs(ctx,path,req,keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown orgstore query endpoint")
		}
	}
}



// nolint: unparam
func queryOrg(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	org := keeper.GetOrg(ctx, path[1])

	res, err := codec.MarshalJSONIndent(keeper.cdc, org)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryOrgs(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var orgsList types.QueryOrgsResolve

	iterator :=keeper.GetOrgIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		store := ctx.KVStore(keeper.storeKey)
		bz := store.Get(iterator.Key())
		var org types.QueryOrgResolve
		keeper.cdc.MustUnmarshalBinaryBare(bz,&org)
		orgsList = append(orgsList, org)
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, orgsList)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
