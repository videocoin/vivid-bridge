package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "vivid-bridge/testutil/keeper"
	"vivid-bridge/testutil/nullify"
	"vivid-bridge/x/vividbridge/keeper"
	"vivid-bridge/x/vividbridge/types"
)

func createNVividGuardians(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VividGuardians {
	items := make([]types.VividGuardians, n)
	for i := range items {
		items[i].Id = keeper.AppendVividGuardians(ctx, items[i])
	}
	return items
}

func TestVividGuardiansGet(t *testing.T) {
	keeper, ctx := keepertest.VividbridgeKeeper(t)
	items := createNVividGuardians(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetVividGuardians(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestVividGuardiansRemove(t *testing.T) {
	keeper, ctx := keepertest.VividbridgeKeeper(t)
	items := createNVividGuardians(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVividGuardians(ctx, item.Id)
		_, found := keeper.GetVividGuardians(ctx, item.Id)
		require.False(t, found)
	}
}

func TestVividGuardiansGetAll(t *testing.T) {
	keeper, ctx := keepertest.VividbridgeKeeper(t)
	items := createNVividGuardians(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVividGuardians(ctx)),
	)
}

func TestVividGuardiansCount(t *testing.T) {
	keeper, ctx := keepertest.VividbridgeKeeper(t)
	items := createNVividGuardians(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetVividGuardiansCount(ctx))
}
