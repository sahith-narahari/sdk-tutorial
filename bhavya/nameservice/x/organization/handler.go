package orgs

import (
	"fmt"
	orgs "github.com/cosmos/bhavya/nameservice/x/organization/internal/keeper"
	"github.com/cosmos/bhavya/nameservice/x/organization/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "Orgs" type messages.
func NewHandler(keeper orgs.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSetOrg:
			return handleMsgSetOrg(ctx, keeper, msg)
		case MsgUpdateOrg:
			return handleMsgUpdateOrg(ctx, keeper, msg)
		case MsgDeleteOrg:
			return handleMsgDeleteOrg(ctx,keeper,msg)

		default:
			errMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle a message to set name
func handleMsgSetOrg(ctx sdk.Context, keeper orgs.Keeper, msg types.MsgSetOrg) sdk.Result {

	org := Orgs{
		Name:  msg.Name,
		CEO:   msg.CEO,
		Owner:msg.Owner,
	}

	keeper.SetOrg(ctx, msg.Name, org) // If so, set the name to the value specified in the msg.
	return sdk.Result{}               // return
}

// Handle a message to set name
func handleMsgUpdateOrg(ctx sdk.Context, keeper orgs.Keeper, msg types.MsgUpdateOrg) sdk.Result {

	keeper.UpdateOrg(ctx, msg.Name, msg.ID, msg.Employee) // Set the name to the value specified in the msg.
	return sdk.Result{}                                   // return
}

func handleMsgDeleteOrg(ctx sdk.Context, keeper orgs.Keeper, msg MsgDeleteOrg) sdk.Result {
	if !keeper.IsOrgPresent(ctx, msg.Name) {
		return types.ErrOrgDoesNotExist(types.DefaultCodespace).Result()
	}
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) {
		return sdk.ErrUnauthorized("Incorrect Owner").Result()
	}
	keeper.DeleteOrg(ctx, msg.Name)
	return sdk.Result{}
}