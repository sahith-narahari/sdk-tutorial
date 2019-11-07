package keeper

import (
	"github.com/comdex-blockchain/x/bank"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/codec"
)

type Keeper struct {
	Coinkeeper bank.Keeper

	StoreKey sdk.StoreKey

	cdc *codec.Codec
}

func (k Keeper) SetName(ctx sdk.Context, Name string, id string) {
	store := ctx.KVStore(k.StoreKey)
	store.Set([]byte(Name), []byte(id))
}
