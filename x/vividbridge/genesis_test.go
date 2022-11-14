package vividbridge_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "vivid-bridge/testutil/keeper"
	"vivid-bridge/testutil/nullify"
	"vivid-bridge/x/vividbridge"
	"vivid-bridge/x/vividbridge/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VividbridgeKeeper(t)
	vividbridge.InitGenesis(ctx, *k, genesisState)
	got := vividbridge.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
