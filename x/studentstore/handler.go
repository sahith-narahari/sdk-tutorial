package studentstore

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgStoreStudent:
			return handleMsgStoreStudent(ctx, keeper, msg)

		default:
			errMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgStoreStudent(ctx sdk.Context, keeper Keeper, msg MsgStoreStudent) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Id)) {
		return sdk.ErrUnauthorized("Incorrect owner").Result()
	}

	keeper.SetStudent(ctx, msg.Name, msg.Id, msg.Owner)
	return sdk.Result{}
}
