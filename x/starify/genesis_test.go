package starify_test

import (
	"testing"

	keepertest "github.com/starify-labs/starify/testutil/keeper"
	"github.com/starify-labs/starify/testutil/nullify"
	"github.com/starify-labs/starify/x/starify"
	"github.com/starify-labs/starify/x/starify/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.StarifyKeeper(t)
	starify.InitGenesis(ctx, *k, genesisState)
	got := starify.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
