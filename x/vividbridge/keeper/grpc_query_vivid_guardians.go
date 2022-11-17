package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"vivid-bridge/x/vividbridge/types"
)

func (k Keeper) VividGuardiansAll(c context.Context, req *types.QueryAllVividGuardiansRequest) (*types.QueryAllVividGuardiansResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var vividGuardianss []types.VividGuardians
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	vividGuardiansStore := prefix.NewStore(store, types.KeyPrefix(types.VividGuardiansKey))

	pageRes, err := query.Paginate(vividGuardiansStore, req.Pagination, func(key []byte, value []byte) error {
		var vividGuardians types.VividGuardians
		if err := k.cdc.Unmarshal(value, &vividGuardians); err != nil {
			return err
		}

		vividGuardianss = append(vividGuardianss, vividGuardians)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVividGuardiansResponse{VividGuardians: vividGuardianss, Pagination: pageRes}, nil
}

func (k Keeper) VividGuardians(c context.Context, req *types.QueryGetVividGuardiansRequest) (*types.QueryGetVividGuardiansResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	vividGuardians, found := k.GetVividGuardians(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetVividGuardiansResponse{VividGuardians: vividGuardians}, nil
}
