package employeestore

import (
"fmt"

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
		keeper.SetInfo(ctx, record.Value, record)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Whois
	iterator := k.GetNamesIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		name := string(iterator.Key())
		whois := k.GetWhois(ctx, name)
		records = append(records, whois)

	}
	return GenesisState{WhoisRecords: records}
}

