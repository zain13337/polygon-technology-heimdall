package simulation

import (
	"time"

	"github.com/maticnetwork/heimdall/milestone/types"
	hmTypes "github.com/maticnetwork/heimdall/types"
	"github.com/maticnetwork/heimdall/types/module"
)

// RandomizedGenState return dummy genesis
func RandomizedGenState(simState *module.SimulationState) {

	startBlock := uint64(0)
	endBlock := uint64(256)
	rootHash := hmTypes.HexToHeimdallHash("123")

	proposerAddress := hmTypes.HexToHeimdallAddress("123")
	timestamp := uint64(time.Now().Unix())
	borChainID := "1234"

	milestone := hmTypes.CreateMilestone(
		startBlock,
		endBlock,
		rootHash,
		proposerAddress,
		borChainID,
		timestamp,
	)

	params := types.DefaultParams()
	genesisState := types.NewGenesisState(
		params,
		&milestone,
	)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(genesisState)
}
