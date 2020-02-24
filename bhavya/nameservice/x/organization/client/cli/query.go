package cli

import (
	"fmt"
	"github.com/cosmos/bhavya/nameservice/x/organization/internal/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	orgsQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the orgs module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	orgsQueryCmd.AddCommand(client.GetCommands(
		GetCmdOrgs(storeKey, cdc),
		GetCmdOrg(storeKey, cdc),
	)...)
	return orgsQueryCmd
}

// GetCmdNames queries a list of all names
func GetCmdOrgs(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-orgs",
		Short: "getting orgs",
		// Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/orgs", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get query orgs\n", err)
				return nil
			}

			var out types.QueryResOrgs
			cdc.MustUnmarshalJSON(res, &out)

			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdNames queries a list of all names
func GetCmdOrg(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-org [name]",
		Short: "getting org",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			name := args[0]
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/org/%s", queryRoute, name), nil)
			if err != nil {
				fmt.Printf("could not get query org\n", err)
				return nil
			}

			var out types.QueryResOrg
			cdc.MustUnmarshalJSON(res, &out)

			return cliCtx.PrintOutput(out)
		},
	}
}
