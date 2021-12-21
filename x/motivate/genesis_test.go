package motivate_test

import (
	"testing"

	keepertest "github.com/kilianstallz/motivate/testutil/keeper"
	"github.com/kilianstallz/motivate/testutil/nullify"
	"github.com/kilianstallz/motivate/x/motivate"
	"github.com/kilianstallz/motivate/x/motivate/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MotivateKeeper(t)
	motivate.InitGenesis(ctx, *k, genesisState)
	got := motivate.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
