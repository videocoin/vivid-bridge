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

		VividGuardiansList: []types.VividGuardians{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		VividGuardiansCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VividbridgeKeeper(t)
	vividbridge.InitGenesis(ctx, *k, genesisState)
	got := vividbridge.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.VividGuardiansList, got.VividGuardiansList)
	require.Equal(t, genesisState.VividGuardiansCount, got.VividGuardiansCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
