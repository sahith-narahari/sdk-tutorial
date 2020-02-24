package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/cosmos/cosmos-sdk/client/context"
	"nameservice/x/orgstore/internal/types"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	nameserviceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the orgstore module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	nameserviceQueryCmd.AddCommand(client.GetCommands(
		GetCmdOrgIs(storeKey,cdc),
		GetCmdOrgs(storeKey,cdc),
		GetCmdUsers(storeKey,cdc),
	)...)
	return nameserviceQueryCmd
}


func GetCmdOrgIs(queryRoute string,cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "org [organization-name]",
		Short: "Query organization",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			name := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/org/%s", queryRoute,name), nil)
			if err != nil {
				fmt.Printf("could not resolve whois - %s \n", name)
				return nil
			}

			var out types.QueryOrgResolve
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)

		},
	}
}

func GetCmdOrgs(queryRoute string,cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "orgs",
		Short: "Query all organizations",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/orgs", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not resolve orgs- err %v\n",err)
				return nil
			}

			var out types.QueryOrgsResolve
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)

		},
	}
}

func GetCmdUsers(queryRoute string,cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "emps",
		Short: "Query all employees in organizations",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			name := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/org/%s", queryRoute,name), nil)
			if err != nil {
				fmt.Printf("could not resolve orgs- err %v\n",err)
				return nil
			}

			var out types.QueryOrgResolve
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out.Employees)

		},
	}
}


