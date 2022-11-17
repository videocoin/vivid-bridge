package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "vivid-bridge/testutil/keeper"
	"vivid-bridge/testutil/nullify"
	"vivid-bridge/x/vividbridge/types"
)

func TestVividGuardiansQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.VividbridgeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNVividGuardians(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetVividGuardiansRequest
		response *types.QueryGetVividGuardiansResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetVividGuardiansRequest{Id: msgs[0].Id},
			response: &types.QueryGetVividGuardiansResponse{VividGuardians: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetVividGuardiansRequest{Id: msgs[1].Id},
			response: &types.QueryGetVividGuardiansResponse{VividGuardians: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetVividGuardiansRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.VividGuardians(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestVividGuardiansQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.VividbridgeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNVividGuardians(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllVividGuardiansRequest {
		return &types.QueryAllVividGuardiansRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.VividGuardiansAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.VividGuardians), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.VividGuardians),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.VividGuardiansAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.VividGuardians), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.VividGuardians),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.VividGuardiansAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.VividGuardians),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.VividGuardiansAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
