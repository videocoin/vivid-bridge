package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"vivid-bridge/x/vividbridge/types"
)

func TestVividGuardiansMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateVividGuardians(ctx, &types.MsgCreateVividGuardians{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestVividGuardiansMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateVividGuardians
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateVividGuardians{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateVividGuardians{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateVividGuardians{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateVividGuardians(ctx, &types.MsgCreateVividGuardians{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateVividGuardians(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestVividGuardiansMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteVividGuardians
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteVividGuardians{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteVividGuardians{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteVividGuardians{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateVividGuardians(ctx, &types.MsgCreateVividGuardians{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteVividGuardians(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
