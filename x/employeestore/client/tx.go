package client

import (
	"github.com/comdex-blockchain/client/context"
	"github.com/cosmos-cg-key-management/cosmos-sdk/client"
	"github.com/cosmos-cg-key-management/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sahith-narahari/sdk-tutorial/x/employeestore/internal/types"
	"github.com/spf13/cobra"
)

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	employeeTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Employee store transaction sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	employeeTxCmd.AddCommand(client.PostCommands(
		GetCmdSetName(cdc),
	))
	return employeeTxCmd
}

func GetCmdSetName(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-employee [name] [id]",
		Short: "Set the id of an employee",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI

			msg := types.NewMsgStoreEmployee(args[0], args[1])

			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, true)
		},
	}
}


