package orgs

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	OrgsRecords []Orgs `json:"whois_records"`
}

func NewGenesisState(orgsRecords []Orgs) GenesisState {
	return GenesisState{OrgsRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.OrgsRecords {
		if record.Owner == nil {
			return fmt.Errorf("invalid WhoisRecord: Value: %s. Error: Missing Owner", record.Name)
		}
		if record.Name == "" {
			return fmt.Errorf("invalid WhoisRecord: Owner: %s. Error: Missing Value", record.CEO)
		}
		if record.CEO == "" {
			return fmt.Errorf("invalid WhoisRecord: Value: %s. Error: Missing Price", record.Owner)
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		OrgsRecords: []Orgs{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.OrgsRecords {
		keeper.SetOrg(ctx, record.Name, record)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Orgs
	iterator := k.GetOrgsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		name := string(iterator.Key())
		whois := k.GetOrg(ctx, name)
		records = append(records, whois)

	}
	return GenesisState{OrgsRecords: records}
}
