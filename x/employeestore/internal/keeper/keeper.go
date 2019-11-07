package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/sdk-tutorials/nameservice/x/employeestore/internal/types"
	"log"
)

type Keeper struct{
	storeKey sdk.StoreKey
	cdc *codec.Codec
}

func NewKeeper (storekey sdk.StoreKey, cdc *codec.Codec) Keeper{
	return Keeper{
		storeKey:storekey,
		cdc:cdc,
	}
}

func(k Keeper) InsertEmployeeInfo (ctx sdk.Context, EmployeeId int64, employee types.EmployeeInfo){
	store := ctx.KVStore(k.storeKey)
	bytesArray := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytesArray, uint64(EmployeeId))
	if store.Has(bytesArray){
		log.Println("Employee Id already present")
		return
	}
	store.Set(bytesArray,k.cdc.MustMarshalBinaryBare(employee))
}

func(k Keeper) QueryEmployee (ctx sdk.Context, EmployeeId int64)  types.EmployeeInfo{
	store := ctx.KVStore(k.storeKey)
	bytesArray := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytesArray, uint64(EmployeeId))
	if !store.Has(bytesArray){
		log.Println("Employee does not exist")
		return types.EmployeeInfo{EmployeeName:""}
	}
	empInfo := store.Get(bytesArray)
	var Employee types.EmployeeInfo
	k.cdc.MustUnmarshalBinaryBare(empInfo, &Employee)
	return Employee
}

func (k Keeper) SetInfo(ctx sdk.Context, Id int64, name string) {
	empInfo := k.QueryEmployee(ctx, Id)
	empInfo.EmployeeName = name
	k.InsertEmployeeInfo(ctx, Id, empInfo)
}

//func (k Keeper) GetInfo (ctx sdk.Context, Id int64) types.EmployeeInfo{
//
//}
