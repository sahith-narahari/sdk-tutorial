package keeper

import (
	"fmt"
	"github.com/cosmos/bhavya/nameservice/x/organization/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the nameservice Querier
const (
	QueryOrgs = "orgs"
	QueryOrg  = "org"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryOrgs:
			return queryOrgs(ctx, req, keeper)
		case QueryOrg:
			return queryOrg(ctx, path, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown org query endpoint")
		}
	}
}

func queryOrgs(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var orgsList types.QueryResOrgs

	iterator := keeper.GetOrgsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		fmt.Println("iterator", iterator)
		store := ctx.KVStore(keeper.storeKey)
		bz := store.Get(iterator.Key())
		var org types.QueryResOrg
		keeper.cdc.MustUnmarshalBinaryBare(bz, &org)
		orgsList = append(orgsList, org)
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, orgsList)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryOrg(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {

	org := keeper.GetOrg(ctx, path[1])

	res, err := codec.MarshalJSONIndent(keeper.cdc, org)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
