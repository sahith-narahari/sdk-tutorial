package orgstore

import (
	"fmt"
	"nameservice/x/orgstore/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "orgstore" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case OrgSet:
			return handleOrgSet(ctx,keeper,msg)
		case UserSet:
			return handleUserSet(ctx,keeper,msg)
		case MsgDelete:
			return handleMsgDeleteName(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized orgstore Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}



func handleOrgSet(ctx sdk.Context, keeper Keeper, msg OrgSet) sdk.Result {
	if msg.CompanyName == keeper.GetOrgName(ctx, msg.CompanyName) { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Company Already exists").Result() // If not, throw an error
	}
	keeper.SetOrg(ctx, msg.CompanyName,msg.Address) // If so, set the name to the value specified in the msg.
	return sdk.Result{}                      // return
}

func handleUserSet(ctx sdk.Context, keeper Keeper, msg UserSet) sdk.Result {
	if msg.Name == keeper.GetUserName(ctx, msg.Name,msg.Organization).Name { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Employee Already exists").Result() // If not, throw an error
	}
	keeper.SetUser(ctx, msg.Name,msg.Address,msg.Organization) // If so, set the name to the value specified in the msg.
	return sdk.Result{}                      // return
}

func handleMsgDeleteName(ctx sdk.Context, keeper Keeper, msg MsgDelete) sdk.Result {
	if !keeper.IsNamePresent(ctx, msg.Name) {
		return types.ErrNameDoesNotExist(types.DefaultCodespace).Result()
	}

	keeper.DeleteOrg(ctx, msg.Name)
	return sdk.Result{}
}

