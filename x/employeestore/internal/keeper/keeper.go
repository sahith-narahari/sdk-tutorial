package keeper

import (
	"fmt"
	"github.com/comdex-blockchain/x/bank"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sahith-narahari/sdk-tutorial/x/employeestore/internal/types"
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

func (k Keeper) GetEmployee(ctx sdk.Context, employeeId string) types.MsgEmployee {
	store := ctx.KVStore(k.StoreKey)
	bytes := []byte(employeeId)
	if !store.Has(bytes) {
		fmt.Println("Employee does not exist")
		return types.MsgEmployee{Name: ""}
	}

	info := store.Get(bytes)
	var emp types.MsgEmployee
	k.cdc.MustUnmarshalBinaryBare(info, &emp)
	return emp
}
