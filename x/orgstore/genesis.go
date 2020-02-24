package orgstore

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	OrgRecords []Org `json:"org_records"`
}

func NewGenesisState(orgRecords []Org) GenesisState {
	return GenesisState{OrgRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.OrgRecords {
		if record.Name == "" {
			return fmt.Errorf("invalid Org Record: Value: %s. Error: Missing Owner", record.Name)
		}
		if record.Address == "" {
			return fmt.Errorf("invalid Org Record: Owner: %s. Error: Missing Value", record.Address)
		}
		//if record.Employees == nil {
		//	return fmt.Errorf("invalid WhoisRecord: Value: %s. Error: Missing Price", record.Value)
		//}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		OrgRecords: []Org{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.OrgRecords {
		keeper.SetOrg(ctx, record.Name, record.Address)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Org
	iterator := k.GetOrgIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		name := string(iterator.Key())
		whois := k.GetOrg(ctx, name)
		records = append(records, whois)

	}
	return GenesisState{OrgRecords: records}
}

