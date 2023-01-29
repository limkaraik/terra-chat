package keeper

import (
	"context"

	"chat/x/chat/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetReceived(goCtx context.Context, req *types.QueryGetReceivedRequest) (*types.QueryGetReceivedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var texts []types.Text
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	textStore := prefix.NewStore(store, types.KeyPrefix(types.TextKey))
	receiver := req.Receiver

	pageRes, err := query.Paginate(textStore, req.Pagination, func(key []byte, value []byte) error {
		var text types.Text
		if err := k.cdc.Unmarshal(value, &text); err != nil {
			return err
		}
		if text.Receiver == receiver {
			texts = append(texts, text)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetReceivedResponse{Text: texts, Pagination: pageRes}, nil
}
