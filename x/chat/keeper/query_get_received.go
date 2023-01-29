package keeper

import (
	"context"

	"chat/x/chat/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetReceived(goCtx context.Context, req *types.QueryGetReceivedRequest) (*types.QueryGetReceivedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryGetReceivedResponse{}, nil
}
