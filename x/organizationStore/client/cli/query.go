package cli

import (
	"github.com/spf13/cobra"
	"github.com/sahith-narahari/sdk-tutorial/x/organizationStore/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"fmt"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	nameserviceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the organization module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	nameserviceQueryCmd.AddCommand(client.GetCommands(

		GetCmdOrganizations(storeKey, cdc),
	)...)
	return nameserviceQueryCmd
}


func GetCmdOrganizations(organizationName string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "organizations",
		Short: "orgs",
		// Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("/organizations/%s", organizationName),nil)
			if err != nil {
				fmt.Printf("could not get query organizations\n")
				return nil
			}

			var out types.MsgOrgStore
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}