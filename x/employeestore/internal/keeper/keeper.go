package keeper

import (
	"fmt"
	"github.com/comdex-blockchain/x/bank"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/sahith-narahari/sdk-tutorial/x/employeestore/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	Coinkeeper bank.Keeper

	StoreKey sdk.StoreKey

	cdc *codec.Codec
}

func (k Keeper) SetName(ctx sdk.Context, Name string, id string, empInfo types.MsgStoreEmployee) {
	if empInfo.Name != "" || empInfo.EmployeeId != "" {
		return
	}
	store := ctx.KVStore(k.StoreKey)
	store.Set([]byte(Name), []byte(id))
}

func (k Keeper) GetEmployee(ctx sdk.Context, employeeId string) types.MsgStoreEmployee {
	store := ctx.KVStore(k.StoreKey)
	empBytesData := []byte(employeeId)
	if !store.Has(empBytesData) {
		fmt.Println("Employee does not exist")
		return types.MsgStoreEmployee{Name: ""}
	}

	info := store.Get(empBytesData)
	var emp types.MsgStoreEmployee

	k.cdc.MustMarshalBinaryBare(info, &emp)
	return emp
}
