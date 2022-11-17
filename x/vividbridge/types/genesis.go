package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		VividGuardiansList: []VividGuardians{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in vividGuardians
	vividGuardiansIdMap := make(map[uint64]bool)
	vividGuardiansCount := gs.GetVividGuardiansCount()
	for _, elem := range gs.VividGuardiansList {
		if _, ok := vividGuardiansIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for vividGuardians")
		}
		if elem.Id >= vividGuardiansCount {
			return fmt.Errorf("vividGuardians id should be lower or equal than the last id")
		}
		vividGuardiansIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
