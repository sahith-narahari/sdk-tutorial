package employeestore

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/sdk-tutorials/nameservice/x/employeestore/internal/types"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgSetInfo:
			return handleMsgSetInfo(ctx, keeper, msg)
		default:
			return sdk.ErrUnknownRequest("Unrecognized msg type").Result()
		}
	}
}

func handleMsgSetInfo(ctx sdk.Context, keeper Keeper, msg types.MsgSetInfo) sdk.Result {
	keeper.InsertEmployeeInfo(ctx, msg.EmployeeInfo)
	return sdk.Result{}
}
