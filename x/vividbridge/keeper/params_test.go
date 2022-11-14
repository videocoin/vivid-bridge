package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "vivid-bridge/testutil/keeper"
	"vivid-bridge/x/vividbridge/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.VividbridgeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
