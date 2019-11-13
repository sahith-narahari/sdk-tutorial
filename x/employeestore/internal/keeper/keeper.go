package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/sdk-tutorials/nameservice/x/employeestore/internal/types"
	"log"
	"math/rand"
	"time"
)

type Keeper struct {
	storeKey sdk.StoreKey
	Cdc      *codec.Codec
}

func NewKeeper(storekey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		storeKey: storekey,
		Cdc:      cdc,
	}
}

func (k Keeper) InsertEmployeeInfo(ctx sdk.Context, employee types.EmployeeInfo) {
	store := ctx.KVStore(k.storeKey)
	rand.Seed(time.Now().UnixNano())
	EmployeeId := rand.Intn(20)
	bytesArray := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytesArray, uint64(EmployeeId))
	if store.Has(bytesArray) {
		log.Println("Employee Id already present")
		return
	}
	store.Set(bytesArray, k.Cdc.MustMarshalBinaryBare(employee.EmployeeName))
}

func (k Keeper) QueryEmployee(ctx sdk.Context, EmployeeId int64) types.EmployeeInfo {
	store := ctx.KVStore(k.storeKey)
	bytesArray := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytesArray, uint64(EmployeeId))
	if !store.Has(bytesArray) {
		log.Println("Employee does not exist")
		return types.EmployeeInfo{EmployeeName: ""}
	}
	empInfo := store.Get(bytesArray)
	var Employee types.EmployeeInfo
	k.Cdc.MustUnmarshalBinaryBare(empInfo, &Employee)
	return Employee
}

func (k Keeper) SetInfo(ctx sdk.Context, Id int64, name string) {
	empInfo := k.QueryEmployee(ctx, Id)
	empInfo.EmployeeName = name
	k.InsertEmployeeInfo(ctx, empInfo)
}

func (k Keeper) GetInfo(ctx sdk.Context, name string) types.EmployeeInfo {
	store := ctx.KVStore(k.storeKey)
	bytesArray := []byte(name)
	if !store.Has(bytesArray) {
		log.Println("Employee does not exist")
		return types.EmployeeInfo{}
	}
	empInfo := store.Get(bytesArray)
	var Employee types.EmployeeInfo
	k.Cdc.MustUnmarshalBinaryBare(empInfo, &Employee)
	return Employee
}

func (k Keeper) GetIdIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}
