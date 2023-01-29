package keeper

import (
	"context"

	"chat/x/chat/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendText(goCtx context.Context, msg *types.MsgSendText) (*types.MsgSendTextResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSendTextResponse{}, nil
}
