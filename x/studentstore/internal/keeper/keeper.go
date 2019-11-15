package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/sahith-narahari/sdk-tutorials/x/studentstore/internal/types"
)

type Keeper struct {
	Storekey sdk.StoreKey
	cdc      *codec.Codec
}

func NewKeeper(coinKeeper bank.Keeper, storekey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		storekey,
		cdc,
	}
}

func (k Keeper) SetStudent(ctx sdk.Context, Name string, id string, owner sdk.AccAddress) {
	var student types.MsgStudentStore

	if owner.Empty() {
		return
	}

	student.Name = Name
	student.Id = id
	student.Owner = owner

	store := ctx.KVStore(k.Storekey)
	store.Set([]byte(id), k.cdc.MustMarshalBinaryBare(student))
}

func (k Keeper) GetOwner(ctx sdk.Context, id string) sdk.AccAddress {
	var studentStore types.MsgStudentStore

	store := ctx.KVStore(k.Storekey)
	if !k.IsNamePresent(ctx, id) {
		return nil
	}

	bz := store.Get([]byte(id))

	k.cdc.MustUnmarshalBinaryBare(bz, &studentStore)
	return studentStore.Owner
}

// Check if the name is present in the store or not
func (k Keeper) IsNamePresent(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.Storekey)
	return store.Has([]byte(name))
}
