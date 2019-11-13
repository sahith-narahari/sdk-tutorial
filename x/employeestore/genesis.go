package employeestore

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	EmpInfoRecords []EmployeeInfo `json:"emp_info_records"`
}

func NewGenesisState(whoIsRecords []EmployeeInfo) GenesisState {
	return GenesisState{EmpInfoRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.EmpInfoRecords {
		if record.EmployeeName == "" {
			return fmt.Errorf("invalid EmpInfoRecords: Error: Missing Name")
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		EmpInfoRecords: []EmployeeInfo{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.EmpInfoRecords {
		keeper.InsertEmployeeInfo(ctx, record)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []EmployeeInfo
	iterator := k.GetIdIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		KeyId := string(iterator.Key())
		Id, err := strconv.Atoi(KeyId)
		if err != nil {
			fmt.Println("Unable to convert string to int")
		}
		emp := k.QueryEmployee(ctx, int64(Id))
		records = append(records, emp)

	}
	return GenesisState{EmpInfoRecords: records}
}
