package cli

import (
	"github.com/gogo/protobuf/codec"
	"github.com/spf13/cobra"
	"github.com/cosmos/sdk-tutorials/nameservice/x/employeestore/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/x/auth"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command  {
	empStoreTxCmd := &cobra.Command{
		Use: types.ModuleName,
		Short: "Emp store transaction subcommands",
		DisableFlagParsing: true,
		SuggestionsMinimumDistance: 2,
		RunE: client.ValidateCmd,
	}

	empStoreTxCmd.AddCommand(client.PostCommands(
		GetCmdSetName(cdc),
		)...)

	return empStoreTxCmd
}

func GetCmdSetName(cdc *codec.Codec) *cobra.Command  {
	return &cobra.Command{
		Use: "set-name [name] [value]",
		Short: "set the value associated with a name that you own",
		Args : cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgSetName(args[0], args[1], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
