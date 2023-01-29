package keeper

import (
	"context"

	"chat/x/chat/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendText(goCtx context.Context, msg *types.MsgSendText) (*types.MsgSendTextResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var post = types.Text{
		Sender:   msg.Creator,
		Receiver: msg.Receiver,
		Body:     msg.Body,
	}
	id := k.AppendText(
		ctx,
		post,
	)

	return &types.MsgSendTextResponse{Id: id}, nil
}
