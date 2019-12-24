package organizationStore

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgStoreOrganization:
			return handleMsgStoreOrg(ctx, keeper, msg)
		case MsgStoreOrgUser:
			return handleMsgStoreOrgUser(ctx, keeper, msg)
		case MsgDeleteOrganizations:
			return handleMsgDeleteOrganizations(ctx,keeper,msg)
		case MsgDeleteOrgUser:
			return handleMsgDeleteOrgUser(ctx,keeper,msg)

		default:
			errMsg := fmt.Sprintf("Unrecognized organizationStore Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgDeleteOrgUser(ctx sdk.Context, keeper Keeper, msg MsgDeleteOrgUser) sdk.Result{
	if !keeper.IsNamePresent(ctx, msg.Name) {
		return sdk.ErrUnknownRequest("This organization doesn't exists").Result()
	}
	if !msg.Address.Equals(keeper.GetOwner(ctx, msg.Name)) {
		return sdk.ErrUnauthorized("Incorrect Owner").Result()
	}

	keeper.DeleteOrganizationUser(ctx, msg.OrgName,msg.Name)
	return sdk.Result{}
}

func handleMsgDeleteOrganizations(ctx sdk.Context, keeper Keeper, msg MsgDeleteOrganizations) sdk.Result{
	if !keeper.IsNamePresent(ctx, msg.Name) {
		return sdk.ErrUnknownRequest("This organization doesn't exists").Result()
	}
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) {
		return sdk.ErrUnauthorized("Incorrect Owner").Result()
	}

	keeper.DeleteOrganization(ctx, msg.Name)
	return sdk.Result{}
}

func handleMsgStoreOrg(ctx sdk.Context, keeper Keeper, msg MsgStoreOrganization) sdk.Result {
	if keeper.GetOrganization(ctx, msg.Name).Name != "" {
		return sdk.ErrUnknownRequest("This organization already exists").Result()
	}

	keeper.SetOrganization(ctx, msg.Name, msg.Owner)
	return sdk.Result{}
}

func handleMsgStoreOrgUser(ctx sdk.Context, keeper Keeper, msg MsgStoreOrgUser) sdk.Result {
	if keeper.GetOrganization(ctx, msg.OrgName).Name == " " {
		return sdk.ErrUnknownRequest("This organization doesn't exists").Result()
	}

	if msg.Name == "" {
		return sdk.ErrUnknownRequest("Got empty user data").Result()
	} else {
		keeper.SetOrganizationUser(ctx, msg.OrgName, msg.Address, msg.Role, msg.Name)
	}

	return sdk.Result{}
}
