package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the employee Querier
const (
	QueryEmployeeName = "employee_name"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryEmployeeName:
			return queryEmployeeName(ctx, path[1], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown employee query endpoint")
		}
	}
}

func queryEmployeeName(ctx sdk.Context, path string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	emp := keeper.GetInfo(ctx, path)
	res, err := codec.MarshalJSONIndent(keeper.Cdc, emp)
	if err != nil {
		panic("could not marshal result to JSON")
	}
	return res, nil

}
