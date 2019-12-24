package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/sahith-narahari/sdk-tutorial/x/organizationStore/internal/types"
	"github.com/spf13/cobra"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	orgStoreTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Organization store",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	orgStoreTxCmd.AddCommand(client.PostCommands(
		CmdToStoreOrg(cdc),
		CmdToStoreOrgUser(cdc),
		CmdToDeleteOrganization(cdc),
		CmdToDeleteOrgUser(cdc),
	)...)
	return orgStoreTxCmd
}

func CmdToDeleteOrgUser(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete_org_user [orgName] [userName]",
		Short: "Give user name with is associated with orgName",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteOrgUser(args[0], cliCtx.GetFromAddress(), args[1])

			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func CmdToDeleteOrganization(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete_organization [orgName] ",
		Short: "Give organization name to delete",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteOrganization(args[0], cliCtx.GetFromAddress())

			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func CmdToStoreOrgUser(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set_org_user [orgName] [userName] [userRole]",
		Short: "Set the user details userName and role assosiated with orgName",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgOrganizationUserStore(args[0], cliCtx.GetFromAddress(), args[1], args[2])

			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func CmdToStoreOrg(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set_org_store [name] ",
		Short: "Set the organizaton name",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgOrganizationStore(args[0], cliCtx.GetFromAddress(), args[1], args[2], args[3])

			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
