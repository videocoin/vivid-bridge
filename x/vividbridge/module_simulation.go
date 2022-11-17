package vividbridge

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"vivid-bridge/testutil/sample"
	vividbridgesimulation "vivid-bridge/x/vividbridge/simulation"
	"vivid-bridge/x/vividbridge/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = vividbridgesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateVividGuardians = "op_weight_msg_vivid_guardians"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateVividGuardians int = 100

	opWeightMsgUpdateVividGuardians = "op_weight_msg_vivid_guardians"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateVividGuardians int = 100

	opWeightMsgDeleteVividGuardians = "op_weight_msg_vivid_guardians"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteVividGuardians int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	vividbridgeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		VividGuardiansList: []types.VividGuardians{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		VividGuardiansCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&vividbridgeGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateVividGuardians int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateVividGuardians, &weightMsgCreateVividGuardians, nil,
		func(_ *rand.Rand) {
			weightMsgCreateVividGuardians = defaultWeightMsgCreateVividGuardians
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateVividGuardians,
		vividbridgesimulation.SimulateMsgCreateVividGuardians(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateVividGuardians int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateVividGuardians, &weightMsgUpdateVividGuardians, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateVividGuardians = defaultWeightMsgUpdateVividGuardians
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateVividGuardians,
		vividbridgesimulation.SimulateMsgUpdateVividGuardians(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteVividGuardians int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteVividGuardians, &weightMsgDeleteVividGuardians, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteVividGuardians = defaultWeightMsgDeleteVividGuardians
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteVividGuardians,
		vividbridgesimulation.SimulateMsgDeleteVividGuardians(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
