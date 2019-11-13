package employeestore

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgStoreName:
			return handleMsgStoreName(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized request: %v ", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgStoreName(ctx sdk.Context, keeper Keeper, msg MsgStoreName) sdk.Result {

	empInfo := keeper.GetEmployee(ctx, msg.EmployeeId)

	keeper.SetName(ctx, msg.Name, msg.EmployeeId, empInfo)
	return sdk.Result{}
}
