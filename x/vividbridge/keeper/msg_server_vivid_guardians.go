package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"vivid-bridge/x/vividbridge/types"
)

func (k msgServer) CreateVividGuardians(goCtx context.Context, msg *types.MsgCreateVividGuardians) (*types.MsgCreateVividGuardiansResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var vividGuardians = types.VividGuardians{
		Creator:        msg.Creator,
		Keys:           msg.Keys,
		ExpirationTime: msg.ExpirationTime,
	}

	id := k.AppendVividGuardians(
		ctx,
		vividGuardians,
	)

	return &types.MsgCreateVividGuardiansResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateVividGuardians(goCtx context.Context, msg *types.MsgUpdateVividGuardians) (*types.MsgUpdateVividGuardiansResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var vividGuardians = types.VividGuardians{
		Creator:        msg.Creator,
		Id:             msg.Id,
		Keys:           msg.Keys,
		ExpirationTime: msg.ExpirationTime,
	}

	// Checks that the element exists
	val, found := k.GetVividGuardians(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetVividGuardians(ctx, vividGuardians)

	return &types.MsgUpdateVividGuardiansResponse{}, nil
}

func (k msgServer) DeleteVividGuardians(goCtx context.Context, msg *types.MsgDeleteVividGuardians) (*types.MsgDeleteVividGuardiansResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetVividGuardians(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveVividGuardians(ctx, msg.Id)

	return &types.MsgDeleteVividGuardiansResponse{}, nil
}
